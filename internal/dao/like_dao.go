package dao

import (
	"context"
	"sponge/internal/model"
	"sponge/pkg/global"

	"gorm.io/gorm"
)

type LikeDao struct {
	db *gorm.DB
}

func NewLikeDao() *LikeDao {
	return &LikeDao{
		db: global.DB,
	}
}

// CreateLike 创建点赞记录
func (d *LikeDao) CreateLike(ctx context.Context, like *model.Like) error {
	return d.db.WithContext(ctx).Create(like).Error
}

// DeleteLike 取消点赞（软删除）
func (d *LikeDao) DeleteLike(ctx context.Context, userID, videoID int64) error {
	return d.db.WithContext(ctx).
		Where("user_id = ? AND video_id = ?", userID, videoID).
		Delete(&model.Like{}).Error
}

// GetLike 查询点赞记录
func (d *LikeDao) GetLike(ctx context.Context, userID, videoID int64) (*model.Like, error) {
	var like model.Like
	err := d.db.WithContext(ctx).
		Where("user_id = ? AND video_id = ?", userID, videoID).
		First(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}

// IsLiked 检查是否点赞
func (d *LikeDao) IsLiked(ctx context.Context, userID, videoID int64) bool {
	var count int64
	d.db.WithContext(ctx).Model(&model.Like{}).
		Where("user_id = ? AND video_id = ? AND status = ?", userID, videoID, 1).
		Count(&count)
	return count > 0
}

// GetLikeList 获取用户点赞的视频列表
func (d *LikeDao) GetLikeList(ctx context.Context, userID int64, page, size int) ([]*model.Like, int64, error) {
	var likes []*model.Like
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Like{}).
		Where("user_id = ? AND status = ?", userID, 1)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err := query.Order("created_at DESC").Offset(offset).Limit(size).Find(&likes).Error
	return likes, total, err
}

// GetLikedVideoIDs 批量获取已点赞的视频ID（性能优化）
func (d *LikeDao) GetLikedVideoIDs(ctx context.Context, userID int64, videoIDs []int64) ([]int64, error) {
	if len(videoIDs) == 0 {
		return []int64{}, nil
	}

	var likes []model.Like
	err := d.db.WithContext(ctx).
		Where("user_id = ? AND video_id IN ? AND status = ?", userID, videoIDs, 1).
		Find(&likes).Error
	if err != nil {
		return nil, err
	}

	result := make([]int64, 0, len(likes))
	for _, like := range likes {
		result = append(result, like.VideoID)
	}
	return result, nil
}

// GetVideoLikeCount 获取视频点赞数
func (d *LikeDao) GetVideoLikeCount(ctx context.Context, videoID int64) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.Like{}).
		Where("video_id = ? AND status = ?", videoID, 1).
		Count(&count).Error
	return count, err
}
