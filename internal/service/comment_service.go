package service

import (
	"context"
	"errors"
	"sponge/internal/dao"
	"sponge/internal/model"
)

type CommentService struct {
	commentDao *dao.CommentDao
	videoDao   *dao.VideoDao
}

func NewCommentService(commentDao *dao.CommentDao, videoDao *dao.VideoDao) *CommentService {
	return &CommentService{
		commentDao: commentDao,
		videoDao:   videoDao,
	}
}

// CommentAction 发布/删除评论
func (s *CommentService) CommentAction(ctx context.Context, userID, videoID int64, action int, commentID int64, content string, rootID, parentID int64) (*model.Comment, error) {
	// 检查视频是否存在
	if !s.videoDao.CheckVideoExists(ctx, videoID) {
		return nil, errors.New("视频不存在")
	}

	if action == 1 { // 发布评论
		if content == "" {
			return nil, errors.New("评论内容不能为空")
		}

		comment := &model.Comment{
			UserID:    userID,
			VideoID:   videoID,
			Content:   content,
			LikeCount: 0,
		}

		// 如果是回复评论
		if parentID > 0 {
			parentComment, err := s.commentDao.GetCommentByID(ctx, parentID)
			if err != nil {
				return nil, errors.New("被回复的评论不存在")
			}

			// 设置 root_id 和 parent_id
			if parentComment.RootID == 0 {
				// 回复顶层评论
				comment.RootID = parentComment.ID
			} else {
				// 回复回复（嵌套回复）
				comment.RootID = parentComment.RootID
			}
			comment.ParentID = parentID
		} else {
			// 顶层评论
			comment.RootID = 0
			comment.ParentID = 0
		}

		if err := s.commentDao.CreateComment(ctx, comment); err != nil {
			return nil, err
		}

		// 增加视频评论数（异步）
		go func() {
			_ = s.videoDao.IncrementCommentCount(context.Background(), videoID)
		}()

		return comment, nil

	} else if action == 2 { // 删除评论
		if commentID <= 0 {
			return nil, errors.New("评论ID不能为空")
		}

		// 检查评论是否存在且属于当前用户
		comment, err := s.commentDao.GetCommentByID(ctx, commentID)
		if err != nil {
			return nil, errors.New("评论不存在")
		}

		if comment.UserID != userID {
			return nil, errors.New("无权删除该评论")
		}

		if err := s.commentDao.DeleteComment(ctx, commentID); err != nil {
			return nil, err
		}

		// 减少视频评论数（异步）
		go func() {
			_ = s.videoDao.DecrementCommentCount(context.Background(), videoID)
		}()

		return comment, nil
	}

	return nil, errors.New("无效的操作类型")
}

// GetCommentList 获取评论列表
func (s *CommentService) GetCommentList(ctx context.Context, videoID int64, page, size int) ([]*model.Comment, int64, error) {
	if size <= 0 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}

	comments, total, err := s.commentDao.GetCommentList(ctx, videoID, page, size)
	if err != nil {
		return nil, 0, err
	}

	// 为每个顶层评论加载回复（最多3条）
	for _, comment := range comments {
		replies, _ := s.commentDao.GetReplyList(ctx, comment.ID, 3)
		// 这里可以设置到 comment 的某个字段，但 model 中没有，所以暂时不处理
		_ = replies
	}

	return comments, total, nil
}
