package dto

// FollowActionReq 关注操作请求
type FollowActionReq struct {
	ToUserID int64 `json:"to_user_id,string" binding:"required"`
	Action   int   `json:"action" binding:"required,oneof=1 2"` // 1-关注, 2-取消关注
}

// FollowListReq 关注列表请求
type FollowListReq struct {
	UserID int64 `form:"user_id" json:"user_id,string" binding:"required"`
	Page   int   `form:"page" binding:"min=1"`
	Size   int   `form:"size" binding:"min=1,max=100"`
}

// FollowerListReq 粉丝列表请求
type FollowerListReq struct {
	UserID int64 `form:"user_id" json:"user_id,string" binding:"required"`
	Page   int   `form:"page" binding:"min=1"`
	Size   int   `form:"size" binding:"min=1,max=100"`
}

// FriendListReq 好友列表请求（互关）
type FriendListReq struct {
	Page int `form:"page" binding:"min=1"`
	Size int `form:"size" binding:"min=1,max=100"`
}

// UserRelationRes 用户关系响应（包含关注状态）
type UserRelationRes struct {
	*UserInfoRes
	IsFollow       bool `json:"is_follow"`        // 当前用户是否关注该用户
	IsFollowed     bool `json:"is_followed"`      // 该用户是否关注当前用户
	IsMutualFollow bool `json:"is_mutual_follow"` // 是否互关
}
