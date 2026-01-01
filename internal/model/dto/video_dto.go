package dto

import "time"

// PublishVideoReq 发布视频请求
type PublishVideoReq struct {
	Title       string  `json:"title" binding:"required,max=255"`
	Description string  `json:"description" binding:"max=1000"`
	PlayUrl     string  `json:"play_url" binding:"required,url"`
	CoverUrl    string  `json:"cover_url" binding:"required,url"`
	Width       int     `json:"width" binding:"min=0"`
	Height      int     `json:"height" binding:"min=0"`
	Duration    float32 `json:"duration" binding:"min=0"`
}

// VideoListReq 视频列表请求
type VideoListReq struct {
	UserID int64 `form:"user_id" json:"user_id,string"`
	Page   int   `form:"page" binding:"min=1"`
	Size   int   `form:"size" binding:"min=1,max=100"`
}

// VideoFeedReq 视频流请求
type VideoFeedReq struct {
	LatestTime int64 `form:"latest_time" json:"latest_time,string"` // 可选，用于分页
	Page       int   `form:"page" binding:"min=1"`
	Size       int   `form:"size" binding:"min=1,max=100"`
}

// VideoInfoRes 视频信息响应
type VideoInfoRes struct {
	ID            int64        `json:"id,string"`
	UserID        int64        `json:"user_id,string"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	PlayUrl       string       `json:"play_url"`
	CoverUrl      string       `json:"cover_url"`
	Width         int          `json:"width"`
	Height        int          `json:"height"`
	Duration      float32      `json:"duration"`
	FavoriteCount int          `json:"favorite_count"`
	CommentCount  int          `json:"comment_count"`
	ViewCount     int          `json:"view_count"`
	IsFavorite    bool         `json:"is_favorite"` // 当前用户是否点赞
	PublishTime   time.Time    `json:"publish_time"`
	Author        *UserInfoRes `json:"author"` // 作者信息
}

// PublishVideoResp 发布视频响应
type PublishVideoResp struct {
	VideoID int64 `json:"video_id,string"`
}
