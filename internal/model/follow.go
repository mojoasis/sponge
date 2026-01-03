package model

import "sponge/pkg/constants"

// Follow 粉丝关注表
type Follow struct {
	BaseModel
	UserID     int64 `gorm:"uniqueIndex:uk_user_follower;uniqueIndex:uk_follower_user;not null" json:"user_id,string"`
	FollowerID int64 `gorm:"uniqueIndex:uk_user_follower;uniqueIndex:uk_follower_user;not null" json:"follower_id,string"`
	IsMutual   int8  `gorm:"default:0" json:"is_mutual"` // 1-互关, 0-单向
}

func (Follow) TableName() string { return constants.FollowsTableName }
