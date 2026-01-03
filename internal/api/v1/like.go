package v1

import (
	"sponge/internal/model/dto"
	"sponge/internal/service"
	"sponge/pkg/global"
	"sponge/pkg/res"
	"sponge/pkg/utils"
	"sponge/pkg/xerror"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LikeApi struct {
	likeService  *service.LikeService
	videoService *service.VideoService
	userService  *service.UserService
}

func NewLikeApi(likeService *service.LikeService, videoService *service.VideoService, userService *service.UserService) *LikeApi {
	return &LikeApi{
		likeService:  likeService,
		videoService: videoService,
		userService:  userService,
	}
}

// LikeAction 点赞/取消点赞
// @Summary      点赞/取消点赞操作
// @Description  点赞或取消点赞视频
// @Tags         点赞模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        data  body      dto.LikeActionReq  true  "点赞信息"
// @Success      200   {object}  res.Response "操作成功"
// @Router       /gateway/like/action [post]
func (a *LikeApi) LikeAction(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt64 := userID.(int64)

	var req dto.LikeActionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if err := a.likeService.LikeAction(c.Request.Context(), userIDInt64, req.VideoID, req.Action); err != nil {
		global.Logger.Error("点赞操作失败", zap.Error(err))
		res.FailMsg(err.Error(), c)
		return
	}

	res.Ok(c)
}

// GetLikeList 获取点赞列表
// @Summary      获取点赞列表
// @Description  获取用户点赞的视频列表
// @Tags         点赞模块
// @Accept       json
// @Produce      json
// @Param        user_id  query     int64  true  "用户ID"
// @Param        page     query     int    false "页码" default(1)
// @Param        size     query     int    false "每页数量" default(30)
// @Success      200      {object}  res.Response{data=[]dto.VideoInfoRes} "查询成功"
// @Router       /gateway/like/list [get]
func (a *LikeApi) GetLikeList(c *gin.Context) {
	var req dto.LikeListReq
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

	if userID, exists := c.Get("userID"); exists {
		_ = userID.(int64)
	}

	videoIDs, total, err := a.likeService.GetLikeList(c.Request.Context(), req.UserID, req.Page, req.Size)
	if err != nil {
		global.Logger.Error("获取点赞列表失败", zap.Error(err))
		res.FailMsg("获取点赞列表失败", c)
		return
	}

	if len(videoIDs) == 0 {
		res.OkData(gin.H{
			"list":  []dto.VideoInfoRes{},
			"total": total,
		}, c)
		return
	}

	// 批量查询视频信息（性能优化）
	// 这里需要 VideoService 提供批量查询方法，暂时使用循环
	videoList := make([]dto.VideoInfoRes, 0, len(videoIDs))
	userIDs := make([]int64, 0)
	videoMap := make(map[int64]*dto.VideoInfoRes)

	for _, videoID := range videoIDs {
		video, err := a.videoService.GetVideoByID(c.Request.Context(), videoID)
		if err != nil {
			continue
		}

		var videoInfo dto.VideoInfoRes
		if err := utils.ToDTO(video, &videoInfo); err != nil {
			continue
		}

		videoInfo.IsFavorite = true // 点赞列表中的视频都是已点赞的
		videoMap[videoID] = &videoInfo
		userIDs = append(userIDs, video.UserID)
	}

	// 批量查询用户信息
	users, _ := a.userService.GetUsersByIDs(c.Request.Context(), userIDs)
	userMap := make(map[int64]*dto.UserInfoRes)
	for _, user := range users {
		var userInfo dto.UserInfoRes
		if err := utils.ToDTO(user, &userInfo); err == nil {
			userMap[user.ID] = &userInfo
		}
	}

	// 组装响应
	for _, videoID := range videoIDs {
		if videoInfo, ok := videoMap[videoID]; ok {
			if author, ok := userMap[videoInfo.UserID]; ok {
				videoInfo.Author = author
			}
			videoList = append(videoList, *videoInfo)
		}
	}

	res.OkData(gin.H{
		"list":  videoList,
		"total": total,
	}, c)
}
