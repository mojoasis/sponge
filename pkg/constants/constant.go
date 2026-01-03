package constants

// 表名常量
const (
	UserTableName        = "users"
	FavoritesTableName   = "likes"
	VideosTableName      = "videos"
	FollowsTableName     = "follows"
	CommentTableName     = "comments"
	ChatSessionTableName = "chat_sessions"
	ChatMessageTableName = "chat_messages"
)

// 业务常量
const (
	// 视频相关
	VideoFeedCount       = 30 // 视频流每次返回的数量
	FavoriteActionType   = 1  // 点赞操作
	UnFavoriteActionType = 2  // 取消点赞操作

	// OSS 相关
	VideoBucket  = "oss-screen"
	OssImgBucket = "portal"

	// 测试数据
	TestSign       = "测试账号！ offer"
	TestAva        = "avatar/test1.jpg"
	TestBackground = "background/test1.png"
)

// Redis Key 前缀
const (
	RedisKeyLoginToken = "login:token:" // 登录 Token 前缀，完整格式: login:token:{userID}
)
