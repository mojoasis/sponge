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
	// 指定更新某些字段
	return d.db.WithContext(ctx).Model(user).Updates(map[string]interface{}{
		"nick_name":        user.NickName,
		"avatar":           user.Avatar,
		"signature":        user.Signature,
		"background_image": user.BackgroundImage,
		"email":            user.Email,
	}).Error
}

// GetUserByID 根据 ID 查询
func (d *UserDao) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsersByIDs 根据用户ID列表批量查询用户
func (d *UserDao) GetUsersByIDs(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	var users []*model.User
	err := d.db.WithContext(ctx).Where("id IN ?", userIDs).Find(&users).Error
	return users, err
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

// UpdateColumn 通用的单字段更新方法 (适合修改密码、手机号、状态等)
func (d *UserDao) UpdateColumn(ctx context.Context, uid int64, column string, value interface{}) error {
	// 使用 map 更新可以触发 GORM 的 UpdatedAt 自动更新
	return d.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", uid).
		Update(column, value).Error
}

// CheckPhoneExists 检查手机号是否存在
func (d *UserDao) CheckPhoneExists(ctx context.Context, phone string) (bool, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.User{}).
		Where("phone = ?", phone).
		Count(&count).Error
	return count > 0, err
}
