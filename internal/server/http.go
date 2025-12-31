package server

import (
	"sponge/internal/middleware"
	"sponge/internal/router"
	"sponge/pkg/global"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Engine *gin.Engine
}

func NewHttpServer(rg *router.RouterGroup) *HttpServer {
	r := gin.New()

	// 注入 Zap 日志、Recovery 和跨域
	r.Use(middleware.ZapLogger(global.Logger))
	r.Use(middleware.ZapRecovery(global.Logger, true))
	r.Use(middleware.Cors())

	// 执行路由初始化
	rg.Init(r)

	return &HttpServer{Engine: r}
}
