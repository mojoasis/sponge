package dto

import (
	"mime/multipart"
	"time"
)

// PublishVideoReq 发布视频请求
// PublishVideoReq 发布视频请求
type PublishVideoReq struct {
	Title       string                  `form:"title" json:"title" binding:"required,max=255"`
	Description string                  `form:"description" json:"description" binding:"max=1000"`
	Files       []*multipart.FileHeader `form:"files" binding:"required"`
}

// VideoListReq 视频列表请求
type VideoListReq struct {
	UserID int64 `form:"userId" json:"userId,string"`
	Page   int   `form:"page" binding:"min=1"`
	Size   int   `form:"size" binding:"min=1,max=100"`
}

// VideoFeedReq 视频流请求
type VideoFeedReq struct {
	LatestTime int64 `form:"latestTime" json:"latestTime,string"` // 可选，用于分页
	Page       int   `form:"page" binding:"min=1"`
	Size       int   `form:"size" binding:"min=1,max=100"`
}

// VideoInfoRes 视频信息响应
type VideoInfoRes struct {
	ID            int64        `json:"id,string"`
	UserID        int64        `json:"userId,string"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	PlayUrl       string       `json:"playUrl"`
	CoverUrl      string       `json:"coverUrl"`
	Width         int          `json:"width"`
	Height        int          `json:"height"`
	Duration      float32      `json:"duration"`
	FavoriteCount int          `json:"favoriteCount"`
	CommentCount  int          `json:"commentCount"`
	ViewCount     int          `json:"viewCount"`
	IsFavorite    bool         `json:"isFavorite"` // 当前用户是否点赞
	PublishTime   time.Time    `json:"publishTime"`
	Author        *UserInfoRes `json:"author"` // 作者信息
}

// PublishVideoResp 发布视频响应
type PublishVideoResp struct {
	VideoID int64 `json:"videoId,string"`
}
