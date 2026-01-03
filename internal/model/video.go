package model

import (
	"sponge/pkg/constants"
	"time"
)

// Video 视频表
type Video struct {
	BaseModel
	UserID        int64     `gorm:"index;not null" json:"user_id,string"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Description   string    `gorm:"type:text" json:"description"`
	PlayUrl       string    `gorm:"type:varchar(255);not null" json:"play_url"`
	CoverUrl      string    `gorm:"type:varchar(255);not null" json:"cover_url"`
	Width         int       `json:"width"`
	Height        int       `json:"height"`
	Duration      float32   `json:"duration"`
	FavoriteCount int       `json:"favorite_count"`
	CommentCount  int       `json:"comment_count"`
	ViewCount     int       `json:"view_count"`
	Status        int8      `gorm:"default:1" json:"status"`
	PublishTime   time.Time `json:"publish_time"`
}

func (Video) TableName() string { return constants.VideosTableName }
