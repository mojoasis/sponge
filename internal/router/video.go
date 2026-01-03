package router

import (
	v1 "sponge/internal/api/v1"
	"sponge/internal/middleware"

	"github.com/gin-gonic/gin"
)

type VideoRouterRegistrar struct {
	api *v1.VideoApi
}

func NewVideoRouterRegistrar(api *v1.VideoApi) *VideoRouterRegistrar {
	return &VideoRouterRegistrar{api: api}
}

func (r *VideoRouterRegistrar) RegisterRoutes(rg *gin.RouterGroup) {
	video := rg.Group("/video/v1")
	{
		// 公开路由
		public := video.Group("/")
		{
			public.GET("/feed", r.api.GetVideoFeed)
			public.GET("/list", r.api.GetVideoList)
		}

		// 需要鉴权的路由
		auth := video.Group("/")
		auth.Use(middleware.JWTAuth())
		{
			auth.POST("/publish", r.api.PublishVideo)
		}
	}
}
