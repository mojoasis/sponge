package global

import (
	"sponge/internal/conf"
	"sponge/pkg/xoss"
	"sponge/pkg/xpay"
	"sponge/pkg/xredis"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

var (
	Config       *conf.AppConfig     // 全局配置句柄
	DB           *gorm.DB            // 全局 MySQL 实例
	Redis        *xredis.RedisClient // 全ify后的 Redis 实例
	OSS          *xoss.OSSClient     // 全局 OSS 实例
	S3Client     *s3.Client          // 底层 S3 客户端
	VideoStorage *xoss.OSSClient     // 专门负责视频逻辑的封装类
	ImgStorage   *xoss.OSSClient     // 专门负责图片逻辑的封装类

	// AliPay WxPay  支付
	AliPay xpay.Payer
	WxPay  xpay.Payer
)
