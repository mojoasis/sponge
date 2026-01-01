package v1

import "github.com/google/wire"

// ProviderSet 是 API v1 层的依赖注入集合
var ProviderSet = wire.NewSet(
	NewUserApi,
	NewVideoApi,
	NewFollowApi,
	NewLikeApi,
	NewCommentApi,
	NewChatApi,
)
