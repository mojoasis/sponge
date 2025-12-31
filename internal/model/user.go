package model

import (
	"sponge/internal/consts"
)

// User 状态常量定义
const (
	UserStatusNormal  int8 = 1 // 正常
	UserStatusBanned  int8 = 2 // 禁用
	UserStatusDeleted int8 = 3 // 已注销
)

// User 用户表
type User struct {
	BaseModel
	UserName        string `gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名" json:"user_name"`
	Password        string `gorm:"type:varchar(100);not null;comment:加密后的密码" json:"-"`
	NickName        string `gorm:"type:varchar(50);comment:昵称" json:"nick_name"`
	Phone           string `gorm:"type:varchar(20);uniqueIndex;comment:手机号" json:"phone"`
	Email           string `gorm:"type:varchar(100);uniqueIndex;comment:邮箱" json:"email"`
	Avatar          string `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	Signature       string `gorm:"type:varchar(255);comment:个人简介" json:"signature"`
	BackgroundImage string `gorm:"type:varchar(255);comment:背景图" json:"background_image"`
	Status          int8   `gorm:"type:tinyint;default:1;comment:状态: 1-正常, 2-禁用, 3-注销" json:"status"`
}

func (User) TableName() string {
	return consts.UserTableName
}
