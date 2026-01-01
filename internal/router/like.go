package router

import (
	v1 "sponge/internal/api/v1"
	"sponge/internal/middleware"

	"github.com/gin-gonic/gin"
)

type LikeRouterRegistrar struct {
	api *v1.LikeApi
}

func NewLikeRouterRegistrar(api *v1.LikeApi) *LikeRouterRegistrar {
	return &LikeRouterRegistrar{api: api}
}

func (r *LikeRouterRegistrar) RegisterRoutes(rg *gin.RouterGroup) {
	like := rg.Group("/like")
	{
		// 公开路由
		//public := like.Group("/")
		//{
		//	public.GET("/list", r.api.GetLikeList)
		//}

		// 需要鉴权的路由
		auth := like.Group("/")
		auth.Use(middleware.JWTAuth())
		{
			auth.POST("/action", r.api.LikeAction)
			auth.GET("/list", r.api.GetLikeList)
		}
	}
}
