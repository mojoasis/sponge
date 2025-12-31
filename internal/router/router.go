package router

import (
	_ "sponge/docs"
	v1 "sponge/internal/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterGroup 路由组容器，由 Wire 自动注入 API 实例
type RouterGroup struct {
	// 使用接口而不是具体类型，提高灵活性
	registrars []RouterRegistrar
}

// NewRouterGroup 构造函数，供 Wire 调用
// 接受所有实现了 RouterRegistrar 接口的模块
func NewRouterGroup(userApi *v1.UserApi) *RouterGroup {
	return &RouterGroup{
		registrars: []RouterRegistrar{
			NewUserRouterRegistrar(userApi),
		},
		// 以后增加模块只需在这里添加：
		// NewVideoRouterRegistrar(videoApi),
	}
}

// Init 核心注册逻辑
func (g *RouterGroup) Init(r *gin.Engine) {
	// 1. 全局非业务路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 2. 统一网关前缀
	gateway := r.Group("/gateway")

	// 3. 自动注册所有模块的路由
	for _, registrar := range g.registrars {
		registrar.RegisterRoutes(gateway)
	}
}
