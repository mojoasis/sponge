package service

import "github.com/google/wire"

// ProviderSet 导出所有业务逻辑相关的构造函数
var ProviderSet = wire.NewSet(
	NewUserService,
	NewVideoService,
	NewFollowService,
	NewLikeService,
	NewCommentService,
	NewChatService,
)
