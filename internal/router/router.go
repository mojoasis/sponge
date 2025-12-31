package router

import (
	_ "sponge/docs"
	v1 "sponge/internal/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// RouterGroup 路由组容器，由 Wire 自动注入 API 实例
type RouterGroup struct {
	UserApi *v1.UserApi
	// VideoApi *v1.VideoApi // 以后增加模块只需在这里添加一行
}

// NewRouterGroup 构造函数，供 Wire 调用
func NewRouterGroup(userApi *v1.UserApi) *RouterGroup {
	return &RouterGroup{
		UserApi: userApi,
	}
}

// Init 核心注册逻辑
func (g *RouterGroup) Init(r *gin.Engine) {
	// 1. 全局非业务路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 2. 统一网关前缀
	gateway := r.Group("/gateway")

	// 3. 挂载各模块
	registerUserRouter(gateway, g.UserApi)
	// registerVideoRouter(gateway, g.VideoApi) // 以后增加只需在此注册
}
