package server

import (
	v1 "sponge/internal/api/v1"
	"sponge/internal/middleware"

	"github.com/gin-gonic/gin"

	_ "sponge/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HttpServer struct {
	Engine *gin.Engine
}

// NewHttpServer 初始化 HTTP 服务
// 注意：这里参数接收了 userApi，Wire 会把初始化好的 API 传进来
func NewHttpServer(userApi *v1.UserApi) *HttpServer {
	r := gin.Default()

	// 注册中间件
	// r.Use(middleware.Cors())
	// 全局中间件
	// r.Use(middleware.Logger(), middleware.Recovery())

	// 注册 Swagger 访问路由
	// 访问路径: http://localhost:18888/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 业务路由分组
	gateway := r.Group("/gateway")
	{
		userGroup := gateway.Group("/user/v1")

		// 直接使用注入进来的 userApi 实例
		userGroup.POST("/login", userApi.Login)
		userGroup.POST("/register", userApi.Register)

		authGroup := userGroup.Group("")
		authGroup.Use(middleware.JWTAuth())
		{
			// authGroup.GET("/profile", userApi.UserProfile)
		}
	}
	return &HttpServer{Engine: r}
}
