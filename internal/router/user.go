package router

import (
	v1 "sponge/internal/api/v1"
	"sponge/internal/middleware"

	"github.com/gin-gonic/gin"
)

// UserRouterRegistrar 用户模块路由注册器
type UserRouterRegistrar struct {
	api *v1.UserApi
}

// NewUserRouterRegistrar 创建用户路由注册器
func NewUserRouterRegistrar(api *v1.UserApi) *UserRouterRegistrar {
	return &UserRouterRegistrar{api: api}
}

// RegisterRoutes 实现 RouterRegistrar 接口
func (r *UserRouterRegistrar) RegisterRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user/v1")
	{
		// 公开路由
		user.POST("/register", r.api.Register)
		user.POST("/login", r.api.Login)

		// 需要鉴权的路由
		auth := user.Group("/")
		auth.Use(middleware.JWTAuth())
		{
			auth.GET("/profile", r.api.UserProfile)
		}
	}
}
