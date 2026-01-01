package dto

// LikeActionReq 点赞操作请求
type LikeActionReq struct {
	VideoID int64 `json:"video_id,string" binding:"required"`
	Action  int   `json:"action" binding:"required,oneof=1 2"` // 1-点赞, 2-取消点赞
}

// LikeListReq 点赞列表请求
type LikeListReq struct {
	UserID int64 `form:"user_id" json:"user_id,string" binding:"required"`
	Page   int   `form:"page" binding:"min=1"`
	Size   int   `form:"size" binding:"min=1,max=100"`
}
