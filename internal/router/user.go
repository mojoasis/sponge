package router

import (
	v1 "sponge/internal/api/v1"
	"sponge/internal/middleware"

	"github.com/gin-gonic/gin"
)

// registerUserRouter 注册用户模块路由
func registerUserRouter(rg *gin.RouterGroup, api *v1.UserApi) {
	user := rg.Group("/user/v1")
	{
		user.POST("/register", api.Register)
		user.POST("/login", api.Login)

		// 需要鉴权的
		auth := user.Group("/")
		auth.Use(middleware.JWTAuth())
		{
			auth.GET("/profile", api.UserProfile)
		}
	}
}
