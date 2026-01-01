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

type CommentApi struct {
	commentService *service.CommentService
	userService    *service.UserService
}

func NewCommentApi(commentService *service.CommentService, userService *service.UserService) *CommentApi {
	return &CommentApi{
		commentService: commentService,
		userService:    userService,
	}
}

// CommentAction 评论操作
// @Summary      评论操作
// @Description  发布或删除评论
// @Tags         评论模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        data  body      dto.CommentActionReq  true  "评论信息"
// @Success      200   {object}  res.Response{data=dto.CommentInfoRes} "操作成功"
// @Router       /gateway/comment/action [post]
func (a *CommentApi) CommentAction(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt64 := userID.(int64)

	var req dto.CommentActionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	comment, err := a.commentService.CommentAction(
		c.Request.Context(),
		userIDInt64,
		req.VideoID,
		req.Action,
		req.CommentID,
		req.Content,
		req.RootID,
		req.ParentID,
	)
	if err != nil {
		global.Logger.Error("评论操作失败", zap.Error(err))
		res.FailMsg(err.Error(), c)
		return
	}

	// 转换为 DTO
	var commentInfo dto.CommentInfoRes
	if err := utils.ToDTO(comment, &commentInfo); err != nil {
		global.Logger.Error("数据转换失败", zap.Error(err))
		res.FailMsg("数据转换失败", c)
		return
	}

	// 获取作者信息
	author, err := a.userService.GetUserProfile(c.Request.Context(), comment.UserID)
	if err == nil {
		var authorInfo dto.UserInfoRes
		if err := utils.ToDTO(author, &authorInfo); err == nil {
			commentInfo.Author = &authorInfo
		}
	}

	res.OkData(commentInfo, c)
}

// GetCommentList 获取评论列表
// @Summary      获取评论列表
// @Description  获取视频的评论列表
// @Tags         评论模块
// @Accept       json
// @Produce      json
// @Param        video_id  query     int64  true  "视频ID"
// @Param        page      query     int    false "页码" default(1)
// @Param        size      query     int    false "每页数量" default(20)
// @Success      200       {object}  res.Response{data=[]dto.CommentInfoRes} "查询成功"
// @Router       /gateway/comment/list [get]
func (a *CommentApi) GetCommentList(c *gin.Context) {
	var req dto.CommentListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if req.VideoID <= 0 {
		res.FailCode(xerror.INVALID_PARAMS, "视频ID不能为空", c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 20
	}

	comments, total, err := a.commentService.GetCommentList(c.Request.Context(), req.VideoID, req.Page, req.Size)
	if err != nil {
		global.Logger.Error("获取评论列表失败", zap.Error(err))
		res.FailMsg("获取评论列表失败", c)
		return
	}

	// 批量查询用户信息（性能优化）
	userIDs := make([]int64, 0, len(comments))
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserID)
	}

	users, _ := a.userService.GetUsersByIDs(c.Request.Context(), userIDs)
	userMap := make(map[int64]*dto.UserInfoRes)
	for _, user := range users {
		var userInfo dto.UserInfoRes
		if err := utils.ToDTO(user, &userInfo); err == nil {
			userMap[user.ID] = &userInfo
		}
	}

	// 转换为 DTO
	commentList := make([]dto.CommentInfoRes, 0, len(comments))
	for _, comment := range comments {
		var commentInfo dto.CommentInfoRes
		if err := utils.ToDTO(comment, &commentInfo); err != nil {
			continue
		}

		if author, ok := userMap[comment.UserID]; ok {
			commentInfo.Author = author
		}

		commentList = append(commentList, commentInfo)
	}

	res.OkData(gin.H{
		"list":  commentList,
		"total": total,
	}, c)
}
