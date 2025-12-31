package model

import (
	"sponge/internal/consts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User 实体类
type User struct {
	ID              string `gorm:"primaryKey;type:char(36);comment:用户ID" json:"id"`
	UserName        string `gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名" json:"user_name"`
	Password        string `gorm:"type:varchar(100);not null;comment:加密后的密码" json:"-"` // json:"-" 也就是不返回密码给前端
	NickName        string `gorm:"type:varchar(50);comment:昵称" json:"nick_name"`
	Avatar          string `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	Signature       string `gorm:"type:varchar(255);comment:个人简介" json:"signature"`
	BackgroundImage string `gorm:"type:varchar(255);comment:背景图" json:"background_image"`
}

func (User) TableName() string {
	return consts.UserTableName
}

// BeforeCreate GORM 钩子：插入前生成 UUID V5
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		// 使用 DNS Namespace 和 UserName 生成 UUID V5
		// 完全随机，改用 uuid.NewString() (即V4)
		u.ID = uuid.NewSHA1(uuid.NameSpaceDNS, []byte(u.UserName)).String()
	}
	return
}
