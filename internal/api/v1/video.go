package v1

import (
	"sponge/internal/model"
	"sponge/internal/model/dto"
	"sponge/internal/service"
	"sponge/pkg/global"
	"sponge/pkg/res"
	"sponge/pkg/utils"
	"sponge/pkg/xerror"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VideoApi struct {
	videoService *service.VideoService
	userService  *service.UserService
	likeService  *service.LikeService
}

func NewVideoApi(videoService *service.VideoService, userService *service.UserService, likeService *service.LikeService) *VideoApi {
	return &VideoApi{
		videoService: videoService,
		userService:  userService,
		likeService:  likeService,
	}
}

// PublishVideo 发布视频
// @Summary      发布视频
// @Description  用户发布视频
// @Tags         视频模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        data  body      dto.PublishVideoReq  true  "视频信息"
// @Success      200   {object}  res.Response{data=dto.PublishVideoResp} "发布成功"
// @Failure      400   {object}  res.Response "参数错误"
// @Router       /gateway/video/publish [post]
func (a *VideoApi) PublishVideo(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt64 := userID.(int64)

	var req dto.PublishVideoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := utils.Translate(err)
		global.Logger.Warn("参数校验失败", zap.String("reason", errMsg))
		res.FailMsg(errMsg, c)
		return
	}

	// 转换为 model
	var videoModel model.Video
	if err := utils.ToDTO(&req, &videoModel); err != nil {
		global.Logger.Error("数据转换失败", zap.Error(err))
		res.FailMsg("数据转换失败", c)
		return
	}

	video, err := a.videoService.PublishVideo(c.Request.Context(), userIDInt64, &videoModel)
	if err != nil {
		global.Logger.Error("发布视频失败", zap.Error(err))
		res.FailMsg(err.Error(), c)
		return
	}

	res.OkData(dto.PublishVideoResp{VideoID: video.ID}, c)
}

// GetVideoList 获取用户视频列表
// @Summary      获取用户视频列表
// @Description  获取指定用户的视频列表
// @Tags         视频模块
// @Accept       json
// @Produce      json
// @Param        user_id  query     int64  true  "用户ID"
// @Param        page     query     int    false "页码" default(1)
// @Param        size     query     int    false "每页数量" default(30)
// @Success      200      {object}  res.Response{data=[]dto.VideoInfoRes} "查询成功"
// @Router       /gateway/video/list [get]
func (a *VideoApi) GetVideoList(c *gin.Context) {
	var req dto.VideoListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if req.UserID <= 0 {
		res.FailCode(xerror.INVALID_PARAMS, "用户ID不能为空", c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 30
	}

	// 获取当前用户ID（如果已登录）
	var currentUserID int64
	if userID, exists := c.Get("userID"); exists {
		currentUserID = userID.(int64)
	}

	videos, total, err := a.videoService.GetVideoList(c.Request.Context(), req.UserID, currentUserID, req.Page, req.Size)
	if err != nil {
		global.Logger.Error("获取视频列表失败", zap.Error(err))
		res.FailMsg("获取视频列表失败", c)
		return
	}

	// 转换为 DTO 并填充作者信息和点赞状态
	videoList := make([]dto.VideoInfoRes, 0, len(videos))
	for _, video := range videos {
		var videoInfo dto.VideoInfoRes
		if err := utils.ToDTO(video, &videoInfo); err != nil {
			continue
		}

		// 获取作者信息
		author, err := a.userService.GetUserProfile(c.Request.Context(), video.UserID)
		if err == nil {
			var authorInfo dto.UserInfoRes
			if err := utils.ToDTO(author, &authorInfo); err == nil {
				videoInfo.Author = &authorInfo
			}
		}

		// 检查是否点赞
		if currentUserID > 0 {
			videoInfo.IsFavorite, _ = a.likeService.IsLiked(c.Request.Context(), currentUserID, video.ID)
		}

		videoList = append(videoList, videoInfo)
	}

	res.OkData(gin.H{
		"list":  videoList,
		"total": total,
	}, c)
}

// GetVideoFeed 获取视频流
// @Summary      获取视频流
// @Description  获取推荐视频流
// @Tags         视频模块
// @Accept       json
// @Produce      json
// @Param        latest_time  query     int64  false "最新时间戳（分页）"
// @Param        page        query     int    false "页码" default(1)
// @Param        size        query     int    false "每页数量" default(30)
// @Success      200         {object}  res.Response{data=[]dto.VideoInfoRes} "查询成功"
// @Router       /gateway/video/feed [get]
func (a *VideoApi) GetVideoFeed(c *gin.Context) {
	var req dto.VideoFeedReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if req.Size <= 0 {
		req.Size = 30
	}

	// 获取当前用户ID（如果已登录）
	var currentUserID int64
	if userID, exists := c.Get("userID"); exists {
		currentUserID = userID.(int64)
	}

	videos, nextTime, err := a.videoService.GetVideoFeed(c.Request.Context(), currentUserID, req.LatestTime, req.Size)
	if err != nil {
		global.Logger.Error("获取视频流失败", zap.Error(err))
		res.FailMsg("获取视频流失败", c)
		return
	}

	// 转换为 DTO
	videoList := make([]dto.VideoInfoRes, 0, len(videos))
	userIDs := make([]int64, 0)
	videoMap := make(map[int64]*model.Video)

	for _, video := range videos {
		videoMap[video.UserID] = video
		userIDs = append(userIDs, video.UserID)
	}

	// 批量查询用户信息（性能优化）
	users, _ := a.userService.GetUsersByIDs(c.Request.Context(), userIDs)
	userMap := make(map[int64]*dto.UserInfoRes)
	for _, user := range users {
		var userInfo dto.UserInfoRes
		if err := utils.ToDTO(user, &userInfo); err == nil {
			userMap[user.ID] = &userInfo
		}
	}

	// 批量查询点赞状态（性能优化）
	var likedVideoIDs []int64
	if currentUserID > 0 && len(videos) > 0 {
		videoIDs := make([]int64, 0, len(videos))
		for _, v := range videos {
			videoIDs = append(videoIDs, v.ID)
		}
		likedVideoIDs, _ = a.likeService.GetLikedVideoIDs(c.Request.Context(), currentUserID, videoIDs)
	}
	likedMap := make(map[int64]bool)
	for _, id := range likedVideoIDs {
		likedMap[id] = true
	}

	// 组装响应
	for _, video := range videos {
		var videoInfo dto.VideoInfoRes
		if err := utils.ToDTO(video, &videoInfo); err != nil {
			continue
		}

		if author, ok := userMap[video.UserID]; ok {
			videoInfo.Author = author
		}

		videoInfo.IsFavorite = likedMap[video.ID]
		videoList = append(videoList, videoInfo)
	}

	res.OkData(gin.H{
		"list":      videoList,
		"next_time": nextTime,
	}, c)
}
