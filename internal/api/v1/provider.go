package v1

import "github.com/google/wire"

// ProviderSet 是 API v1 层的依赖注入集合
// 以后每增加一个模块（如 VideoApi），只需在此 Set 中添加其 New 函数
var ProviderSet = wire.NewSet(
	NewUserApi,
	// NewVideoApi, // 以后取消注释
	// NewChatApi,  // 以后取消注释
)
