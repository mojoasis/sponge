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

type ChatApi struct {
	chatService *service.ChatService
	userService *service.UserService
}

func NewChatApi(chatService *service.ChatService, userService *service.UserService) *ChatApi {
	return &ChatApi{
		chatService: chatService,
		userService: userService,
	}
}

// SendMessage 发送消息
// @Summary      发送消息
// @Description  发送聊天消息
// @Tags         聊天模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        data  body      dto.SendMessageReq  true  "消息信息"
// @Success      200   {object}  res.Response{data=dto.MessageInfoRes} "发送成功"
// @Router       /gateway/chat/message [post]
func (a *ChatApi) SendMessage(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt64 := userID.(int64)

	var req dto.SendMessageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	message, err := a.chatService.SendMessage(c.Request.Context(), userIDInt64, req.ToUserID, req.MsgType, req.Content)
	if err != nil {
		global.Logger.Error("发送消息失败", zap.Error(err))
		res.FailMsg(err.Error(), c)
		return
	}

	// 转换为 DTO
	var messageInfo dto.MessageInfoRes
	if err := utils.ToDTO(message, &messageInfo); err != nil {
		global.Logger.Error("数据转换失败", zap.Error(err))
		res.FailMsg("数据转换失败", c)
		return
	}

	res.OkData(messageInfo, c)
}

// GetMessageList 获取消息列表
// @Summary      获取消息列表
// @Description  获取与指定用户的消息列表
// @Tags         聊天模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        to_user_id  query     int64  true  "对方用户ID"
// @Param        page        query     int    false "页码" default(1)
// @Param        size        query     int    false "每页数量" default(20)
// @Success      200         {object}  res.Response{data=[]dto.MessageInfoRes} "查询成功"
// @Router       /gateway/chat/message/list [get]
func (a *ChatApi) GetMessageList(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt64 := userID.(int64)

	var req dto.MessageListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if req.ToUserID <= 0 {
		res.FailCode(xerror.INVALID_PARAMS, "对方用户ID不能为空", c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 20
	}

	messages, total, err := a.chatService.GetMessageList(c.Request.Context(), userIDInt64, req.ToUserID, req.Page, req.Size)
	if err != nil {
		global.Logger.Error("获取消息列表失败", zap.Error(err))
		res.FailMsg("获取消息列表失败", c)
		return
	}

	// 转换为 DTO
	messageList := make([]dto.MessageInfoRes, 0, len(messages))
	for _, message := range messages {
		var messageInfo dto.MessageInfoRes
		if err := utils.ToDTO(message, &messageInfo); err != nil {
			continue
		}
		messageList = append(messageList, messageInfo)
	}

	res.OkData(gin.H{
		"list":  messageList,
		"total": total,
	}, c)
}

// GetChatSessionList 获取会话列表
// @Summary      获取会话列表
// @Description  获取用户的聊天会话列表
// @Tags         聊天模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        page  query     int  false "页码" default(1)
// @Param        size  query     int  false "每页数量" default(20)
// @Success      200   {object}  res.Response{data=[]dto.ChatSessionRes} "查询成功"
// @Router       /gateway/chat/session/list [get]
func (a *ChatApi) GetChatSessionList(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt64 := userID.(int64)

	var req dto.ChatSessionListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		errMsg := utils.Translate(err)
		res.FailMsg(errMsg, c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 20
	}

	sessions, total, err := a.chatService.GetChatSessionList(c.Request.Context(), userIDInt64, req.Page, req.Size)
	if err != nil {
		global.Logger.Error("获取会话列表失败", zap.Error(err))
		res.FailMsg("获取会话列表失败", c)
		return
	}

	// 批量查询用户信息
	peerIDs := make([]int64, 0, len(sessions))
	for _, session := range sessions {
		peerIDs = append(peerIDs, session.PeerID)
	}

	users, _ := a.userService.GetUsersByIDs(c.Request.Context(), peerIDs)
	userMap := make(map[int64]*dto.UserInfoRes)
	for _, user := range users {
		var userInfo dto.UserInfoRes
		if err := utils.ToDTO(user, &userInfo); err == nil {
			userMap[user.ID] = &userInfo
		}
	}

	// 转换为 DTO
	sessionList := make([]dto.ChatSessionRes, 0, len(sessions))
	for _, session := range sessions {
		var sessionInfo dto.ChatSessionRes
		if err := utils.ToDTO(session, &sessionInfo); err != nil {
			continue
		}

		if peerUser, ok := userMap[session.PeerID]; ok {
			sessionInfo.PeerUser = peerUser
		}

		sessionList = append(sessionList, sessionInfo)
	}

	res.OkData(gin.H{
		"list":  sessionList,
		"total": total,
	}, c)
}
