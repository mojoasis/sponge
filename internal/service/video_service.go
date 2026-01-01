package service

import (
	"context"
	"errors"
	"sponge/internal/consts"
	"sponge/internal/dao"
	"sponge/internal/model"
	"time"

	"gorm.io/gorm"
)

type VideoService struct {
	videoDao *dao.VideoDao
	userDao  *dao.UserDao
	likeDao  *dao.LikeDao
}

func NewVideoService(videoDao *dao.VideoDao, userDao *dao.UserDao, likeDao *dao.LikeDao) *VideoService {
	return &VideoService{
		videoDao: videoDao,
		userDao:  userDao,
		likeDao:  likeDao,
	}
}

// PublishVideo 发布视频
func (s *VideoService) PublishVideo(ctx context.Context, userID int64, req *model.Video) (*model.Video, error) {
	video := &model.Video{
		UserID:        userID,
		Title:         req.Title,
		Description:   req.Description,
		PlayUrl:       req.PlayUrl,
		CoverUrl:      req.CoverUrl,
		Width:         req.Width,
		Height:        req.Height,
		Duration:      req.Duration,
		FavoriteCount: 0,
		CommentCount:  0,
		ViewCount:     0,
		Status:        1,
		PublishTime:   time.Now(),
	}

	if err := s.videoDao.CreateVideo(ctx, video); err != nil {
		return nil, err
	}

	return video, nil
}

// GetVideoList 获取用户视频列表
func (s *VideoService) GetVideoList(ctx context.Context, userID int64, currentUserID int64, page, size int) ([]*model.Video, int64, error) {
	if size <= 0 {
		size = consts.VideoFeedCount
	}
	if page <= 0 {
		page = 1
	}

	videos, total, err := s.videoDao.GetVideosByUserID(ctx, userID, page, size)
	if err != nil {
		return nil, 0, err
	}

	// 批量查询点赞状态（性能优化）
	if currentUserID > 0 && len(videos) > 0 {
		videoIDs := make([]int64, 0, len(videos))
		for _, v := range videos {
			videoIDs = append(videoIDs, v.ID)
		}
		// 这里会在 LikeService 中处理
	}

	return videos, total, nil
}

// GetVideoFeed 获取视频流
func (s *VideoService) GetVideoFeed(ctx context.Context, currentUserID int64, latestTime int64, size int) ([]*model.Video, int64, error) {
	if size <= 0 {
		size = consts.VideoFeedCount
	}

	videos, err := s.videoDao.GetVideoFeed(ctx, latestTime, size)
	if err != nil {
		return nil, 0, err
	}

	var nextTime int64
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreatedAt.UnixMilli()
	}

	return videos, nextTime, nil
}

// GetVideoByID 获取视频详情
func (s *VideoService) GetVideoByID(ctx context.Context, videoID int64, currentUserID int64) (*model.Video, error) {
	video, err := s.videoDao.GetVideoByID(ctx, videoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("视频不存在")
		}
		return nil, err
	}

	// 增加播放数（异步，不阻塞）
	go func() {
		_ = s.videoDao.IncrementViewCount(context.Background(), videoID)
	}()

	return video, nil
}

// CheckVideoExists 检查视频是否存在
func (s *VideoService) CheckVideoExists(ctx context.Context, videoID int64) bool {
	_, err := s.videoDao.GetVideoByID(ctx, videoID)
	return err == nil
}
