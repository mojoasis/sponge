package initialize

import (
	"context"
	"fmt"
	"log"
	"os"
	"sponge/pkg/utils"
	"sponge/pkg/xoss"
	"strings"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"sponge/internal/conf"
	"sponge/pkg/global"
	"sponge/pkg/xredis"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NamingClient 定义一个全局 NamingClient，方便后续可能的注销操作
var NamingClient naming_client.INamingClient

// InitConfig 启动总入口
func InitConfig() {
	// 最先初始化日志，确保后续报错都能被记录
	global.Logger = initLogger()

	// 初始化验证翻译器
	if err := utils.InitTrans(); err != nil {
		global.Logger.Error("初始化翻译器失败", zap.Error(err))
	}

	// 读取本地 config.yaml 获取 Nacos 连接信息
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取本地配置失败: %s", err))
	}

	// 初始化雪花算法
	machineID := v.GetInt64("system.machineID")
	if err := utils.InitSnowflake(machineID); err != nil {
		global.Logger.Fatal("初始化雪花算法失败", zap.Error(err))
	}

	// 初始化 Nacos 客户端
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:   v.GetString("nacos.host"),
			Port:     v.GetUint64("nacos.port"),
			GrpcPort: v.GetUint64("nacos.rpcPort"),
		},
	}
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(v.GetString("nacos.namespace")),
		constant.WithUsername(v.GetString("nacos.user")),
		constant.WithPassword(v.GetString("nacos.pass")),
		constant.WithTimeoutMs(5000),
	)
	// 从 Nacos 获取业务配置
	configClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		panic(fmt.Sprintf("初始化Nacos失败: %v", err))
	}
	// 加载 Nacos 上 yaml 文件的配置
	dataID := v.GetString("nacos.dataId")
	group := v.GetString("nacos.group")
	refreshConfig(configClient, dataID, group)

	// 将当前服务注册到 Nacos
	NamingClient, err = clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		panic(fmt.Errorf("创建 Nacos 服务发现客户端失败: %w", err))
	}

	// 这里开始注册到 Nacos
	success, err := NamingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          utils.GetLocalIP(),                // 当前服务 IP，建议通过工具类自动获取本地 IP
		Port:        uint64(global.Config.Server.Port), // 当前服务端口
		ServiceName: global.Config.Server.Name,         // 当前服务名称
		GroupName:   v.GetString("nacos.group"),        // 默认 DEFAULT_GROUP
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true, // 临时实例，程序挂了 Nacos 会自动剔除
	})

	if !success || err != nil {
		panic(fmt.Errorf("服务注册失败: %v", err))
	}

	fmt.Printf("服务 %s 成功注册到 Nacos\n", v.GetString("server.name"))

	// 注册监听，实现热更新
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("检测到Nacos配置变更，正在热更新实例...")
			refreshConfig(configClient, dataID, group)
		},
	})
}

// refreshConfig 从 Nacos 拉取最新内容并刷新全局变量
func refreshConfig(client config_client.IConfigClient, dataID, group string) {
	content, err := client.GetConfig(vo.ConfigParam{DataId: dataID, Group: group})
	if err != nil {
		fmt.Printf("获取远程配置失败: %v\n", err)
		return
	}

	// 使用 viper 解析 YAML 字符串
	runtimeViper := viper.New()
	runtimeViper.SetConfigType("yaml")
	if err := runtimeViper.ReadConfig(strings.NewReader(content)); err != nil {
		fmt.Printf("解析远程配置失败: %v\n", err)
		return
	}

	var newConf conf.AppConfig
	if err := runtimeViper.Unmarshal(&newConf); err != nil {
		fmt.Printf("配置转换失败: %v\n", err)
		return
	}

	// 更新全局变量句柄
	global.Config = &newConf

	// 执行各个组件的实例化
	initDatabase(&newConf)
	initRedis(&newConf)
	initOSS(&newConf)

	// initPay(&newConf) ... 同理
}

// initDatabase 创建数据库连接
//
//	func initDatabase(cfg *conf.AppConfig) {
//		db, err := gorm.Open(mysql.Open(cfg.MySQL.DSN), &gorm.Config{})
//		if err != nil {
//			fmt.Printf("MySQL 连接失败: %v\n", err)
//			return
//		}
//		global.DB = db
//		fmt.Println("MySQL 初始化成功")
//	}
func initDatabase(cfg *conf.AppConfig) {
	// 如果是热更新直接跳过
	if global.DB != nil {
		return
	}
	// 1. 配置日志与慢 SQL 监控
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值：超过 200ms 则记录为慢 SQL
			LogLevel:                  logger.Info,            // 日志级别：Warn 级别会打印慢 SQL 和 错误
			IgnoreRecordNotFoundError: true,                   // 忽略 ErrRecordNotFound 错误日志
			Colorful:                  true,                   // 彩色打印
		},
	)

	// 2. 配置 GORM 全局映射规则
	gormConfig := &gorm.Config{
		Logger: newLogger, // 注入上面定义的日志配置
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置为 true 则不使用复数，User 对应 user 表
			// TablePrefix: "t_", // 如果有统一的表前缀可以设置
		},
	}
	// 3. 建立连接
	db, err := gorm.Open(mysql.Open(cfg.MySQL.DSN), gormConfig)
	if err != nil {
		log.Printf("MySQL 连接失败: %v\n", err)
		return
	}
	// 4. 设置连接池 (提高性能的关键)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	global.DB = db
}

// initRedis 初始化Redis
func initRedis(cfg *conf.AppConfig) {
	// 	// 如果是热更新直接跳过
	if global.Redis != nil {
		return
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	// 验证连接
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		fmt.Printf("Redis 连接失败: %v\n", err)
		return
	}
	// 赋值给全局封装好的 Redis 工具类
	global.Redis = &xredis.RedisClient{Client: rdb}
	fmt.Println("Redis 初始化成功")
}

// initOSS 初始化 OSS
func initOSS(cfg *conf.AppConfig) {
	//  创建 S3 客户端
	client := xoss.NewOSSClient(
		cfg.OSS.Endpoint,
		cfg.OSS.AccessKey,
		cfg.OSS.SecretKey,
	)
	if client == nil {
		fmt.Println("OSS 客户端初始化失败")
		return
	}
	// 赋值给全局封装好的 OSS 工具类
	global.OSS = &xoss.OSSClient{
		S3Client: client,
	}
	fmt.Println("OSS 初始化成功")
}

// initLogger 初始化日志
func initLogger() *zap.Logger {
	// 生产环境
	//logger, _ := zap.NewProduction()
	//config := zap.NewDevelopmentConfig()
	// 开发环境
	logger, _ := zap.NewDevelopment()
	config := zap.NewDevelopmentConfig()
	// 加上这个，日志里会显示调用者的文件名和行号
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 替换全局的 zap.L()
	zap.ReplaceGlobals(logger)
	return logger
}
