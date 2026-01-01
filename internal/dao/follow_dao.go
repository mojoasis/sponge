package dao

import (
	"context"
	"sponge/internal/model"
	"sponge/pkg/global"

	"gorm.io/gorm"
)

type FollowDao struct {
	db *gorm.DB
}

func NewFollowDao() *FollowDao {
	return &FollowDao{
		db: global.DB,
	}
}

// CreateFollow 创建关注关系
func (d *FollowDao) CreateFollow(ctx context.Context, follow *model.Follow) error {
	return d.db.WithContext(ctx).Create(follow).Error
}

// DeleteFollow 取消关注
func (d *FollowDao) DeleteFollow(ctx context.Context, userID, followerID int64) error {
	return d.db.WithContext(ctx).
		Where("user_id = ? AND follower_id = ?", userID, followerID).
		Delete(&model.Follow{}).Error
}

// GetFollow 查询关注关系
func (d *FollowDao) GetFollow(ctx context.Context, userID, followerID int64) (*model.Follow, error) {
	var follow model.Follow
	err := d.db.WithContext(ctx).
		Where("user_id = ? AND follower_id = ?", userID, followerID).
		First(&follow).Error
	if err != nil {
		return nil, err
	}
	return &follow, nil
}

// IsFollowing 检查是否关注
func (d *FollowDao) IsFollowing(ctx context.Context, userID, followerID int64) bool {
	var count int64
	d.db.WithContext(ctx).Model(&model.Follow{}).
		Where("user_id = ? AND follower_id = ?", userID, followerID).
		Count(&count)
	return count > 0
}

// GetFollowList 获取关注列表（我关注的人）
func (d *FollowDao) GetFollowList(ctx context.Context, userID int64, page, size int) ([]*model.Follow, int64, error) {
	var follows []*model.Follow
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Follow{}).Where("follower_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err := query.Order("created_at DESC").Offset(offset).Limit(size).Find(&follows).Error
	return follows, total, err
}

// GetFollowerList 获取粉丝列表（关注我的人）
func (d *FollowDao) GetFollowerList(ctx context.Context, userID int64, page, size int) ([]*model.Follow, int64, error) {
	var follows []*model.Follow
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Follow{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err := query.Order("created_at DESC").Offset(offset).Limit(size).Find(&follows).Error
	return follows, total, err
}

// GetFriendList 获取好友列表（互关）
func (d *FollowDao) GetFriendList(ctx context.Context, userID int64, page, size int) ([]*model.Follow, int64, error) {
	var follows []*model.Follow
	var total int64

	// 查询互关关系（is_mutual = 1）
	query := d.db.WithContext(ctx).Model(&model.Follow{}).
		Where("follower_id = ? AND is_mutual = ?", userID, 1)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err := query.Order("created_at DESC").Offset(offset).Limit(size).Find(&follows).Error
	return follows, total, err
}

// GetFollowCount 获取关注数
func (d *FollowDao) GetFollowCount(ctx context.Context, userID int64) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.Follow{}).
		Where("follower_id = ?", userID).
		Count(&count).Error
	return count, err
}

// GetFollowerCount 获取粉丝数
func (d *FollowDao) GetFollowerCount(ctx context.Context, userID int64) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.Follow{}).
		Where("user_id = ?", userID).
		Count(&count).Error
	return count, err
}

// BatchCheckFollow 批量检查关注关系（性能优化）
func (d *FollowDao) BatchCheckFollow(ctx context.Context, userID int64, targetUserIDs []int64) (map[int64]bool, error) {
	if len(targetUserIDs) == 0 {
		return make(map[int64]bool), nil
	}

	var follows []model.Follow
	err := d.db.WithContext(ctx).
		Where("follower_id = ? AND user_id IN ?", userID, targetUserIDs).
		Find(&follows).Error
	if err != nil {
		return nil, err
	}

	result := make(map[int64]bool)
	for _, follow := range follows {
		result[follow.UserID] = true
	}
	return result, nil
}

// UpdateMutualStatus 更新互关状态
func (d *FollowDao) UpdateMutualStatus(ctx context.Context, userID, followerID int64, isMutual int8) error {
	// 更新两条记录
	err1 := d.db.WithContext(ctx).Model(&model.Follow{}).
		Where("user_id = ? AND follower_id = ?", userID, followerID).
		Update("is_mutual", isMutual).Error

	err2 := d.db.WithContext(ctx).Model(&model.Follow{}).
		Where("user_id = ? AND follower_id = ?", followerID, userID).
		Update("is_mutual", isMutual).Error

	if err1 != nil {
		return err1
	}
	return err2
}
