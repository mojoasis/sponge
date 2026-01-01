package router

import (
	v1 "sponge/internal/api/v1"
	"sponge/internal/middleware"

	"github.com/gin-gonic/gin"
)

type CommentRouterRegistrar struct {
	api *v1.CommentApi
}

func NewCommentRouterRegistrar(api *v1.CommentApi) *CommentRouterRegistrar {
	return &CommentRouterRegistrar{api: api}
}

func (r *CommentRouterRegistrar) RegisterRoutes(rg *gin.RouterGroup) {
	comment := rg.Group("/comment")
	{
		// 公开路由
		public := comment.Group("/")
		{
			public.GET("/list", r.api.GetCommentList)
		}

		// 需要鉴权的路由
		auth := comment.Group("/")
		auth.Use(middleware.JWTAuth())
		{
			auth.POST("/action", r.api.CommentAction)
		}
	}
}
