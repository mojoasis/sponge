package service

import (
	"context"
	"errors"
	"sponge/internal/dao"
	"sponge/internal/model"
	"sponge/pkg/constants"
)

type FollowService struct {
	followDao *dao.FollowDao
	userDao   *dao.UserDao
}

func NewFollowService(followDao *dao.FollowDao, userDao *dao.UserDao) *FollowService {
	return &FollowService{
		followDao: followDao,
		userDao:   userDao,
	}
}

// FollowAction 关注/取消关注操作
func (s *FollowService) FollowAction(ctx context.Context, userID, toUserID int64, action int) error {
	if userID == toUserID {
		return errors.New("不能关注自己")
	}

	// 检查目标用户是否存在
	_, err := s.userDao.GetUserByID(ctx, toUserID)
	if err != nil {
		return errors.New("目标用户不存在")
	}

	if action == constants.FavoriteActionType { // 1-关注
		// 检查是否已关注
		if s.followDao.IsFollowing(ctx, toUserID, userID) {
			return errors.New("已经关注过了")
		}

		// 创建关注关系
		follow := &model.Follow{
			UserID:     toUserID,
			FollowerID: userID,
			IsMutual:   0,
		}

		if err := s.followDao.CreateFollow(ctx, follow); err != nil {
			return err
		}

		// 检查是否互关（对方也关注了我）
		if s.followDao.IsFollowing(ctx, userID, toUserID) {
			// 更新两条记录的互关状态
			_ = s.followDao.UpdateMutualStatus(ctx, userID, toUserID, 1)
		}

	} else if action == constants.UnFavoriteActionType { // 2-取消关注
		// 检查是否已关注
		if !s.followDao.IsFollowing(ctx, toUserID, userID) {
			return errors.New("未关注该用户")
		}

		// 删除关注关系
		if err := s.followDao.DeleteFollow(ctx, toUserID, userID); err != nil {
			return err
		}

		// 如果之前是互关，更新对方记录的互关状态
		if s.followDao.IsFollowing(ctx, userID, toUserID) {
			_ = s.followDao.UpdateMutualStatus(ctx, userID, toUserID, 0)
		}
	}

	return nil
}

// GetFollowList 获取关注列表
func (s *FollowService) GetFollowList(ctx context.Context, userID, currentUserID int64, page, size int) ([]int64, int64, error) {
	if size <= 0 {
		size = 30
	}
	if page <= 0 {
		page = 1
	}

	follows, total, err := s.followDao.GetFollowList(ctx, userID, page, size)
	if err != nil {
		return nil, 0, err
	}

	userIDs := make([]int64, 0, len(follows))
	for _, follow := range follows {
		userIDs = append(userIDs, follow.UserID)
	}

	return userIDs, total, nil
}

// GetFollowerList 获取粉丝列表
func (s *FollowService) GetFollowerList(ctx context.Context, userID, currentUserID int64, page, size int) ([]int64, int64, error) {
	if size <= 0 {
		size = 30
	}
	if page <= 0 {
		page = 1
	}

	follows, total, err := s.followDao.GetFollowerList(ctx, userID, page, size)
	if err != nil {
		return nil, 0, err
	}

	userIDs := make([]int64, 0, len(follows))
	for _, follow := range follows {
		userIDs = append(userIDs, follow.FollowerID)
	}

	return userIDs, total, nil
}

// GetFriendList 获取好友列表（互关）
func (s *FollowService) GetFriendList(ctx context.Context, userID int64, page, size int) ([]int64, int64, error) {
	if size <= 0 {
		size = 30
	}
	if page <= 0 {
		page = 1
	}

	follows, total, err := s.followDao.GetFriendList(ctx, userID, page, size)
	if err != nil {
		return nil, 0, err
	}

	userIDs := make([]int64, 0, len(follows))
	for _, follow := range follows {
		userIDs = append(userIDs, follow.UserID)
	}

	return userIDs, total, nil
}

// IsFollowing 检查是否关注
func (s *FollowService) IsFollowing(ctx context.Context, userID, targetUserID int64) bool {
	return s.followDao.IsFollowing(ctx, targetUserID, userID)
}

// BatchCheckFollow 批量检查关注关系（性能优化）
func (s *FollowService) BatchCheckFollow(ctx context.Context, userID int64, targetUserIDs []int64) (map[int64]bool, error) {
	return s.followDao.BatchCheckFollow(ctx, userID, targetUserIDs)
}
