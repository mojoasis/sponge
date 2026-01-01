package v1

import (
	"sponge/internal/model/dto"
	"sponge/internal/service"
	"sponge/pkg/global"
	"sponge/pkg/res"
	"sponge/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FollowApi struct {
	followService *service.FollowService
	userService   *service.UserService
}

func NewFollowApi(followService *service.FollowService, userService *service.UserService) *FollowApi {
	return &FollowApi{
		followService: followService,
		userService:   userService,
	}
}

// FollowAction 关注/取消关注
// @Summary      关注/取消关注操作
// @Description  关注或取消关注用户
// @Tags         关注模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        data  body      dto.FollowActionReq  true  "关注信息"
// @Success      200   {object}  res.Response "操作成功"
// @Router       /gateway/follow/action [post]
func (a *FollowApi) FollowAction(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt64 := userID.(int64)

	var req dto.FollowActionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if err := a.followService.FollowAction(c.Request.Context(), userIDInt64, req.ToUserID, req.Action); err != nil {
		global.Logger.Error("关注操作失败", zap.Error(err))
		res.FailMsg(err.Error(), c)
		return
	}

	res.Ok(c)
}

// GetFollowList 获取关注列表
// @Summary      获取关注列表
// @Description  获取用户的关注列表
// @Tags         关注模块
// @Accept       json
// @Produce      json
// @Param        user_id  query     int64  true  "用户ID"
// @Param        page     query     int    false "页码" default(1)
// @Param        size     query     int    false "每页数量" default(30)
// @Success      200      {object}  res.Response{data=[]dto.UserRelationRes} "查询成功"
// @Router       /gateway/follow/list [get]
func (a *FollowApi) GetFollowList(c *gin.Context) {
	var req dto.FollowListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 30
	}

	var currentUserID int64
	if userID, exists := c.Get("userID"); exists {
		currentUserID = userID.(int64)
	}

	userIDs, total, err := a.followService.GetFollowList(c.Request.Context(), req.UserID, currentUserID, req.Page, req.Size)
	if err != nil {
		global.Logger.Error("获取关注列表失败", zap.Error(err))
		res.FailMsg("获取关注列表失败", c)
		return
	}

	// 批量查询用户信息
	users, err := a.userService.GetUsersByIDs(c.Request.Context(), userIDs)
	if err != nil {
		global.Logger.Error("获取用户信息失败", zap.Error(err))
		res.FailMsg("获取用户信息失败", c)
		return
	}

	// 批量检查关注关系（性能优化）
	var followMap map[int64]bool
	if currentUserID > 0 {
		followMap, _ = a.followService.BatchCheckFollow(c.Request.Context(), currentUserID, userIDs)
	}

	// 转换为 DTO
	userList := make([]dto.UserRelationRes, 0, len(users))
	for _, user := range users {
		var userInfo dto.UserInfoRes
		if err := utils.ToDTO(user, &userInfo); err != nil {
			continue
		}

		relation := dto.UserRelationRes{
			UserInfoRes: &userInfo,
			IsFollow:    followMap[user.ID],
		}

		// 检查是否被关注（互关）
		if currentUserID > 0 {
			relation.IsFollowed = a.followService.IsFollowing(c.Request.Context(), currentUserID, user.ID)
			relation.IsMutualFollow = relation.IsFollow && relation.IsFollowed
		}

		userList = append(userList, relation)
	}

	res.OkData(gin.H{
		"list":  userList,
		"total": total,
	}, c)
}

// GetFollowerList 获取粉丝列表
// @Summary      获取粉丝列表
// @Description  获取用户的粉丝列表
// @Tags         关注模块
// @Accept       json
// @Produce      json
// @Param        user_id  query     int64  true  "用户ID"
// @Param        page     query     int    false "页码" default(1)
// @Param        size     query     int    false "每页数量" default(30)
// @Success      200      {object}  res.Response{data=[]dto.UserRelationRes} "查询成功"
// @Router       /gateway/follower/list [get]
func (a *FollowApi) GetFollowerList(c *gin.Context) {
	var req dto.FollowerListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 30
	}

	var currentUserID int64
	if userID, exists := c.Get("userID"); exists {
		currentUserID = userID.(int64)
	}

	userIDs, total, err := a.followService.GetFollowerList(c.Request.Context(), req.UserID, currentUserID, req.Page, req.Size)
	if err != nil {
		global.Logger.Error("获取粉丝列表失败", zap.Error(err))
		res.FailMsg("获取粉丝列表失败", c)
		return
	}

	// 批量查询用户信息
	users, err := a.userService.GetUsersByIDs(c.Request.Context(), userIDs)
	if err != nil {
		global.Logger.Error("获取用户信息失败", zap.Error(err))
		res.FailMsg("获取用户信息失败", c)
		return
	}

	// 批量检查关注关系
	var followMap map[int64]bool
	if currentUserID > 0 {
		followMap, _ = a.followService.BatchCheckFollow(c.Request.Context(), currentUserID, userIDs)
	}

	// 转换为 DTO
	userList := make([]dto.UserRelationRes, 0, len(users))
	for _, user := range users {
		var userInfo dto.UserInfoRes
		if err := utils.ToDTO(user, &userInfo); err != nil {
			continue
		}

		relation := dto.UserRelationRes{
			UserInfoRes: &userInfo,
			IsFollow:    followMap[user.ID],
			IsFollowed:  true, // 粉丝列表中的用户都关注了当前用户
		}

		if currentUserID > 0 {
			relation.IsMutualFollow = relation.IsFollow && relation.IsFollowed
		}

		userList = append(userList, relation)
	}

	res.OkData(gin.H{
		"list":  userList,
		"total": total,
	}, c)
}

// GetFriendList 获取好友列表（互关）
// @Summary      获取好友列表
// @Description  获取互相关注的好友列表
// @Tags         关注模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        page  query     int  false "页码" default(1)
// @Param        size  query     int  false "每页数量" default(30)
// @Success      200   {object}  res.Response{data=[]dto.UserRelationRes} "查询成功"
// @Router       /gateway/friend/list [get]
func (a *FollowApi) GetFriendList(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt64 := userID.(int64)

	var req dto.FriendListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 30
	}

	userIDs, total, err := a.followService.GetFriendList(c.Request.Context(), userIDInt64, req.Page, req.Size)
	if err != nil {
		global.Logger.Error("获取好友列表失败", zap.Error(err))
		res.FailMsg("获取好友列表失败", c)
		return
	}

	// 批量查询用户信息
	users, err := a.userService.GetUsersByIDs(c.Request.Context(), userIDs)
	if err != nil {
		global.Logger.Error("获取用户信息失败", zap.Error(err))
		res.FailMsg("获取用户信息失败", c)
		return
	}

	// 转换为 DTO（好友都是互关的）
	userList := make([]dto.UserRelationRes, 0, len(users))
	for _, user := range users {
		var userInfo dto.UserInfoRes
		if err := utils.ToDTO(user, &userInfo); err != nil {
			continue
		}

		relation := dto.UserRelationRes{
			UserInfoRes:    &userInfo,
			IsFollow:       true,
			IsFollowed:     true,
			IsMutualFollow: true,
		}

		userList = append(userList, relation)
	}

	res.OkData(gin.H{
		"list":  userList,
		"total": total,
	}, c)
}
