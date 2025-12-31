package dao

import (
	"context"
	"sponge/internal/model"
	"sponge/pkg/global"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

// NewUserDao 构造函数
// Wire 会调用这个函数来生成 DAO 实例
func NewUserDao() *UserDao {
	return &UserDao{
		db: global.DB, // 在这里绑定全局 DB，以后 DAO 内部只用 d.db
	}
}

// CreateUser 创建用户
func (d *UserDao) CreateUser(ctx context.Context, user *model.User) error {
	return d.db.WithContext(ctx).Create(user).Error
}

// UpdateUserInfo 更新用户信息
func (d *UserDao) UpdateUserInfo(ctx context.Context, user *model.User) error {
	// 指定更新某些字段，防止密码被覆盖
	return d.db.WithContext(ctx).Model(user).Updates(map[string]interface{}{
		"nick_name": user.NickName,
		"avatar":    user.Avatar,
		"signature": user.Signature,
	}).Error
}

// GetUserByID 根据 ID 查询
func (d *UserDao) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}

// GetUserByUsername 根据用户名查询
func (d *UserDao) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
