package dao

import (
	"context"
	"sponge/internal/model"
	"sponge/pkg/global"
	"time"

	"gorm.io/gorm"
)

type VideoDao struct {
	db *gorm.DB
}

func NewVideoDao() *VideoDao {
	return &VideoDao{
		db: global.DB,
	}
}

// CreateVideo 创建视频
func (d *VideoDao) CreateVideo(ctx context.Context, video *model.Video) error {
	return d.db.WithContext(ctx).Create(video).Error
}

// GetVideoByID 根据ID查询视频
func (d *VideoDao) GetVideoByID(ctx context.Context, id int64) (*model.Video, error) {
	var video model.Video
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

// GetVideosByUserID 根据用户ID查询视频列表
func (d *VideoDao) GetVideosByUserID(ctx context.Context, userID int64, page, size int) ([]*model.Video, int64, error) {
	var videos []*model.Video
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Video{}).Where("user_id = ?", userID)

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * size
	err := query.Order("created_at DESC").Offset(offset).Limit(size).Find(&videos).Error
	return videos, total, err
}

// GetVideoFeed 获取视频流（按时间倒序）
func (d *VideoDao) GetVideoFeed(ctx context.Context, latestTime int64, size int) ([]*model.Video, error) {
	var videos []*model.Video
	query := d.db.WithContext(ctx).Model(&model.Video{}).Where("status = ?", 1)

	// 如果有 latest_time，用于分页
	if latestTime > 0 {
		latestTimeObj := time.Unix(latestTime/1000, 0)
		query = query.Where("created_at < ?", latestTimeObj)
	}

	err := query.Order("created_at DESC").Limit(size).Find(&videos).Error
	return videos, err
}

// GetVideosByIDs 批量查询视频（性能优化）
func (d *VideoDao) GetVideosByIDs(ctx context.Context, videoIDs []int64) ([]*model.Video, error) {
	if len(videoIDs) == 0 {
		return []*model.Video{}, nil
	}
	var videos []*model.Video
	err := d.db.WithContext(ctx).Where("id IN ?", videoIDs).Find(&videos).Error
	return videos, err
}

// IncrementFavoriteCount 增加点赞数
func (d *VideoDao) IncrementFavoriteCount(ctx context.Context, videoID int64) error {
	return d.db.WithContext(ctx).Model(&model.Video{}).
		Where("id = ?", videoID).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
}

// DecrementFavoriteCount 减少点赞数
func (d *VideoDao) DecrementFavoriteCount(ctx context.Context, videoID int64) error {
	return d.db.WithContext(ctx).Model(&model.Video{}).
		Where("id = ?", videoID).
		UpdateColumn("favorite_count", gorm.Expr("GREATEST(favorite_count - 1, 0)")).Error
}

// IncrementCommentCount 增加评论数
func (d *VideoDao) IncrementCommentCount(ctx context.Context, videoID int64) error {
	return d.db.WithContext(ctx).Model(&model.Video{}).
		Where("id = ?", videoID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
}

// DecrementCommentCount 减少评论数
func (d *VideoDao) DecrementCommentCount(ctx context.Context, videoID int64) error {
	return d.db.WithContext(ctx).Model(&model.Video{}).
		Where("id = ?", videoID).
		UpdateColumn("comment_count", gorm.Expr("GREATEST(comment_count - 1, 0)")).Error
}

// IncrementViewCount 增加播放数
func (d *VideoDao) IncrementViewCount(ctx context.Context, videoID int64) error {
	return d.db.WithContext(ctx).Model(&model.Video{}).
		Where("id = ?", videoID).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}

// CheckVideoExists 检查视频是否存在
func (d *VideoDao) CheckVideoExists(ctx context.Context, videoID int64) bool {
	var count int64
	d.db.WithContext(ctx).Model(&model.Video{}).
		Where("id = ?", videoID).
		Count(&count)
	return count > 0
}
