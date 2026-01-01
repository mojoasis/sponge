package router

import (
	v1 "sponge/internal/api/v1"
	"sponge/internal/middleware"

	"github.com/gin-gonic/gin"
)

type FollowRouterRegistrar struct {
	api *v1.FollowApi
}

func NewFollowRouterRegistrar(api *v1.FollowApi) *FollowRouterRegistrar {
	return &FollowRouterRegistrar{api: api}
}

func (r *FollowRouterRegistrar) RegisterRoutes(rg *gin.RouterGroup) {
	follow := rg.Group("/follow")
	{
		// 公开路由
		public := follow.Group("/")
		{
			public.GET("/list", r.api.GetFollowList)
		}

		// 需要鉴权的路由
		auth := follow.Group("/")
		auth.Use(middleware.JWTAuth())
		{
			auth.POST("/action", r.api.FollowAction)
		}
	}

	// 粉丝列表
	follower := rg.Group("/follower")
	{
		follower.GET("/list", r.api.GetFollowerList)
	}

	// 好友列表（互关）
	friend := rg.Group("/friend")
	{
		friend.Use(middleware.JWTAuth())
		{
			friend.GET("/list", r.api.GetFriendList)
		}
	}
}
