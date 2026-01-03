package v1

import (
	"fmt"
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

// PublishVideo 上传视频
//
//	@Summary      批量上传视频
//	@Description  采用流式上传技术，在上传过程中实时完成视频抽帧。接口返回视频播放地址和封面图地址。
//	@Tags         视频模块
//	@Security     Bearer
//	@Accept       multipart/form-data
//	@Produce      json
//
// @Param        title        formData  string  true  "视频标题"
// @Param        description  formData  string  false "视频描述"
// @Param        files        formData  file    true  "视频文件(多选)"
//
//	@Success      200 {object} res.Response "发布成功"
//	@Failure      400 {object} res.Response "参数校验失败"
//	@Failure      401 {object} res.Response "登录失效"
//	@Failure      500 {object} res.Response "服务器内部错误"
//	@Router       /gateway/video/v1/publish [post]
func (a *VideoApi) PublishVideo(c *gin.Context) {
	var req dto.PublishVideoReq
	fmt.Println("req:", req)
	if err := c.ShouldBind(&req); err != nil {
		res.FailMsg(utils.Translate(err), c)
		return
	}
	// 获取当前用户ID
	userID := res.GetUserID(c)
	if userID == 0 {
		res.FailMsg("登录失效", c)
		return
	}
	// 用于存储处理结果切片
	err := a.videoService.PublishVideo(c.Request.Context(), userID, &req)
	if err != nil {
		return
	}

	// 返回结果
	res.OkMsg("发布成功", c)
}

// GetVideoList 获取用户视频列表
// @Summary      获取用户视频列表
// @Description  获取指定用户的视频列表
// @Tags         视频模块
// @Accept       json
// @Produce      json
// @Param        userId  query     int64  true  "用户ID"
// @Param        page     query     int    false "页码" default(1)
// @Param        size     query     int    false "每页数量" default(30)
// @Success      200      {object}  res.Response{data=[]dto.VideoInfoRes} "查询成功"
// @Router       /gateway/video/v1/list [get]
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
// @Param        latestTime  query     int64  false "最新时间戳（分页）"
// @Param        page        query     int    false "页码" default(1)
// @Param        size        query     int    false "每页数量" default(30)
// @Success      200         {object}  res.Response{data=[]dto.VideoInfoRes} "查询成功"
// @Router       /gateway/video/v1/feed [get]
func (a *VideoApi) GetVideoFeed(c *gin.Context) {
	var req dto.VideoFeedReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}
	// 分页参数
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
