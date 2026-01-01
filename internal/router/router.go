package router

import (
	_ "sponge/docs"
	v1 "sponge/internal/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/wdcbot/qingfeng"
)

// RouterGroup 路由组容器，由 Wire 自动注入 API 实例
type RouterGroup struct {
	// 使用接口而不是具体类型，提高灵活性
	registrars []RouterRegistrar
}

// NewRouterGroup 构造函数，供 Wire 调用
// 接受所有实现了 RouterRegistrar 接口的模块
func NewRouterGroup(
	userApi *v1.UserApi,
	videoApi *v1.VideoApi,
	followApi *v1.FollowApi,
	likeApi *v1.LikeApi,
	commentApi *v1.CommentApi,
	chatApi *v1.ChatApi,
) *RouterGroup {
	return &RouterGroup{
		registrars: []RouterRegistrar{
			NewUserRouterRegistrar(userApi),
			NewVideoRouterRegistrar(videoApi),
			NewFollowRouterRegistrar(followApi),
			NewLikeRouterRegistrar(likeApi),
			NewCommentRouterRegistrar(commentApi),
			NewChatRouterRegistrar(chatApi),
		},
	}
}

// Init 核心注册逻辑
func (g *RouterGroup) Init(r *gin.Engine) {
	// swagger
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 替换青峰Swag
	r.GET("/doc/*any", qingfeng.Handler(qingfeng.Config{
		Title:    "sponge",
		BasePath: "/doc",
		DocPath:  "./docs/swagger.json",
	}))

	// 2. 统一网关前缀
	gateway := r.Group("/gateway")

	// 3. 自动注册所有模块的路由
	for _, registrar := range g.registrars {
		registrar.RegisterRoutes(gateway)
	}
}
