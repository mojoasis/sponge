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
// 参考 tiktok_demo 的路由分组方式
func (r *UserRouterRegistrar) RegisterRoutes(rg *gin.RouterGroup) {
	// 用户模块路由组
	user := rg.Group("/user/v1")
	{
		// 公开路由（无需鉴权）
		public := user.Group("/")
		{
			public.POST("/register", r.api.Register)
			public.POST("/login", r.api.Login)
		}
		// 需要鉴权的路由
		auth := user.Group("/")
		auth.Use(middleware.JWTAuth())
		{
			auth.GET("/profile", r.api.UserProfile)
			auth.POST("/update/profile", r.api.UpdateUserProfile)
			auth.POST("/update/phone", r.api.UpdatePhone)
			auth.POST("/update/password", r.api.UpdatePassword)
		}
	}
}
