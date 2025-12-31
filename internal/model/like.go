package model

// Like 点赞表
type Like struct {
	BaseModel
	UserID  int64 `gorm:"uniqueIndex:uk_user_video;not null" json:"user_id,string"`
	VideoID int64 `gorm:"uniqueIndex:uk_user_video;index;not null" json:"video_id,string"`
	Status  int8  `gorm:"default:1;comment:1-点赞, 0-取消点赞" json:"status"`
}
