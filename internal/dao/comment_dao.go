package dao

import (
	"context"
	"sponge/internal/model"
	"sponge/pkg/global"

	"gorm.io/gorm"
)

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao() *CommentDao {
	return &CommentDao{
		db: global.DB,
	}
}

// CreateComment 创建评论
func (d *CommentDao) CreateComment(ctx context.Context, comment *model.Comment) error {
	return d.db.WithContext(ctx).Create(comment).Error
}

// DeleteComment 删除评论（软删除）
func (d *CommentDao) DeleteComment(ctx context.Context, commentID int64) error {
	return d.db.WithContext(ctx).Where("id = ?", commentID).Delete(&model.Comment{}).Error
}

// GetCommentByID 根据ID查询评论
func (d *CommentDao) GetCommentByID(ctx context.Context, id int64) (*model.Comment, error) {
	var comment model.Comment
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetCommentList 获取视频的评论列表（顶层评论）
func (d *CommentDao) GetCommentList(ctx context.Context, videoID int64, page, size int) ([]*model.Comment, int64, error) {
	var comments []*model.Comment
	var total int64

	// 只查询顶层评论（root_id = 0 或 root_id = id）
	query := d.db.WithContext(ctx).Model(&model.Comment{}).
		Where("video_id = ? AND (root_id = 0 OR root_id = id)", videoID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err := query.Order("created_at DESC").Offset(offset).Limit(size).Find(&comments).Error
	return comments, total, err
}

// GetReplyList 获取评论的回复列表
func (d *CommentDao) GetReplyList(ctx context.Context, rootID int64, limit int) ([]*model.Comment, error) {
	var replies []*model.Comment
	query := d.db.WithContext(ctx).Model(&model.Comment{}).
		Where("root_id = ? AND id != ?", rootID, rootID)

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Order("created_at ASC").Find(&replies).Error
	return replies, err
}

// GetReplyCount 获取评论的回复数量
func (d *CommentDao) GetReplyCount(ctx context.Context, rootID int64) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.Comment{}).
		Where("root_id = ? AND id != ?", rootID, rootID).
		Count(&count).Error
	return count, err
}

// IncrementLikeCount 增加评论点赞数
func (d *CommentDao) IncrementLikeCount(ctx context.Context, commentID int64) error {
	return d.db.WithContext(ctx).Model(&model.Comment{}).
		Where("id = ?", commentID).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
}

// DecrementLikeCount 减少评论点赞数
func (d *CommentDao) DecrementLikeCount(ctx context.Context, commentID int64) error {
	return d.db.WithContext(ctx).Model(&model.Comment{}).
		Where("id = ?", commentID).
		UpdateColumn("like_count", gorm.Expr("GREATEST(like_count - 1, 0)")).Error
}
