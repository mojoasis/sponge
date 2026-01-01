package model

import "sponge/internal/consts"

// Comment 评论表
type Comment struct {
	BaseModel
	UserID    int64  `gorm:"index;not null" json:"user_id,string"`
	VideoID   int64  `gorm:"index;not null" json:"video_id,string"`
	RootID    int64  `gorm:"index;not null;default:0" json:"root_id,string"` // 顶层评论ID
	ParentID  int64  `gorm:"not null;default:0" json:"parent_id,string"`     // 被回复的评论ID
	Content   string `gorm:"type:text;not null" json:"content"`
	LikeCount int    `gorm:"default:0" json:"like_count"`
}

func (Comment) TableName() string { return consts.CommentTableName }
