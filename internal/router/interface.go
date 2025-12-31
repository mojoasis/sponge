package router

import "github.com/gin-gonic/gin"

// RouterRegistrar 路由注册器接口
// 每个模块可以实现此接口来注册自己的路由
type RouterRegistrar interface {
	RegisterRoutes(rg *gin.RouterGroup)
}
