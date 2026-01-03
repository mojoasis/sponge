package service

import (
	"context"
	"errors"
	"sponge/internal/dao"
	"sponge/internal/model"
	"sponge/pkg/constants"
)

type LikeService struct {
	likeDao  *dao.LikeDao
	videoDao *dao.VideoDao
}

func NewLikeService(likeDao *dao.LikeDao, videoDao *dao.VideoDao) *LikeService {
	return &LikeService{
		likeDao:  likeDao,
		videoDao: videoDao,
	}
}

// LikeAction 点赞/取消点赞操作
func (s *LikeService) LikeAction(ctx context.Context, userID, videoID int64, action int) error {
	// 检查视频是否存在
	if !s.videoDao.CheckVideoExists(ctx, videoID) {
		return errors.New("视频不存在")
	}

	if action == constants.FavoriteActionType { // 1-点赞
		// 检查是否已点赞
		if s.likeDao.IsLiked(ctx, userID, videoID) {
			return errors.New("已经点赞过了")
		}

		// 创建点赞记录
		like := &model.Like{
			UserID:  userID,
			VideoID: videoID,
			Status:  1,
		}

		if err := s.likeDao.CreateLike(ctx, like); err != nil {
			return err
		}

		// 增加视频点赞数（异步，不阻塞）
		go func() {
			_ = s.videoDao.IncrementFavoriteCount(context.Background(), videoID)
		}()

	} else if action == constants.UnFavoriteActionType { // 2-取消点赞
		// 检查是否已点赞
		if !s.likeDao.IsLiked(ctx, userID, videoID) {
			return errors.New("未点赞该视频")
		}

		// 删除点赞记录
		if err := s.likeDao.DeleteLike(ctx, userID, videoID); err != nil {
			return err
		}

		// 减少视频点赞数（异步，不阻塞）
		go func() {
			_ = s.videoDao.DecrementFavoriteCount(context.Background(), videoID)
		}()
	}

	return nil
}

// GetLikeList 获取用户点赞的视频列表
func (s *LikeService) GetLikeList(ctx context.Context, userID int64, page, size int) ([]int64, int64, error) {
	if size <= 0 {
		size = 30
	}
	if page <= 0 {
		page = 1
	}

	likes, total, err := s.likeDao.GetLikeList(ctx, userID, page, size)
	if err != nil {
		return nil, 0, err
	}

	videoIDs := make([]int64, 0, len(likes))
	for _, like := range likes {
		videoIDs = append(videoIDs, like.VideoID)
	}

	return videoIDs, total, nil
}

// IsLiked 检查是否点赞
func (s *LikeService) IsLiked(ctx context.Context, userID, videoID int64) (bool, error) {
	return s.likeDao.IsLiked(ctx, userID, videoID), nil
}

// GetLikedVideoIDs 批量获取已点赞的视频ID（性能优化）
func (s *LikeService) GetLikedVideoIDs(ctx context.Context, userID int64, videoIDs []int64) ([]int64, error) {
	return s.likeDao.GetLikedVideoIDs(ctx, userID, videoIDs)
}
