package dao

import "github.com/google/wire"

// ProviderSet 导出所有数据库相关的构造函数
var ProviderSet = wire.NewSet(
	NewUserDao,
	// NewVideoDao, // 以后有了再加
)
