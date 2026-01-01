package dto

import "time"

// CommentActionReq 评论操作请求
type CommentActionReq struct {
	VideoID   int64  `json:"video_id,string" binding:"required"`
	Action    int    `json:"action" binding:"required,oneof=1 2"` // 1-发布评论, 2-删除评论
	CommentID int64  `json:"comment_id,string"`                   // 删除时必填
	Content   string `json:"content"`                             // 发布时必填
	RootID    int64  `json:"root_id,string"`                      // 顶层评论ID（回复时使用）
	ParentID  int64  `json:"parent_id,string"`                    // 被回复的评论ID（回复时使用）
}

// CommentListReq 评论列表请求
type CommentListReq struct {
	VideoID int64 `form:"video_id" json:"video_id,string" binding:"required"`
	Page    int   `form:"page" binding:"min=1"`
	Size    int   `form:"size" binding:"min=1,max=100"`
}

// CommentInfoRes 评论信息响应
type CommentInfoRes struct {
	ID         int64             `json:"id,string"`
	UserID     int64             `json:"user_id,string"`
	VideoID    int64             `json:"video_id,string"`
	RootID     int64             `json:"root_id,string"`
	ParentID   int64             `json:"parent_id,string"`
	Content    string            `json:"content"`
	LikeCount  int               `json:"like_count"`
	CreatedAt  time.Time         `json:"created_at"`
	Author     *UserInfoRes      `json:"author"`            // 评论作者信息
	IsLiked    bool              `json:"is_liked"`          // 当前用户是否点赞该评论
	ReplyCount int               `json:"reply_count"`       // 回复数量（仅顶层评论有）
	Replies    []*CommentInfoRes `json:"replies,omitempty"` // 回复列表（仅顶层评论有）
}
