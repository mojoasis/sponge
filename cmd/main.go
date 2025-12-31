package main

import (
	"fmt"
	"sponge/pkg/global"
	"sponge/pkg/initialize"
)

// @title           Sponge API
// @version         1.0
// @description     基于 Gin + Gorm + Wire 的企业级后端服务
// @termsOfService  https://github.com/Mojitocean/sponge

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:18888
// @BasePath  /gateway
// @securityDefinitions.oauth2 BearerAuth
// @in header
// @name Authorization

func main() {
	// 1. 初始化所有配置和资源
	initialize.InitConfig()

	srv, cleanup, err := InitApp()
	if err != nil {
		panic(fmt.Sprintf("初始化失败: %v", err))
	}
	defer cleanup()

	// 3. 启动服务
	fmt.Printf("服务启动成功，端口: %d\n", global.Config.Server.Port)
	if err := srv.Engine.Run(fmt.Sprintf(":%d", global.Config.Server.Port)); err != nil {
		panic(err)
	}
}
