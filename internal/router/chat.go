package router

import (
	v1 "sponge/internal/api/v1"
	"sponge/internal/middleware"

	"github.com/gin-gonic/gin"
)

type ChatRouterRegistrar struct {
	api *v1.ChatApi
}

func NewChatRouterRegistrar(api *v1.ChatApi) *ChatRouterRegistrar {
	return &ChatRouterRegistrar{api: api}
}

func (r *ChatRouterRegistrar) RegisterRoutes(rg *gin.RouterGroup) {
	chat := rg.Group("/chat")
	chat.Use(middleware.JWTAuth()) // 聊天模块全部需要鉴权
	{
		chat.POST("/message", r.api.SendMessage)
		chat.GET("/message/list", r.api.GetMessageList)
		chat.GET("/session/list", r.api.GetChatSessionList)
	}
}
