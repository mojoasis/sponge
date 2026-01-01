package service

import (
	"context"
	"errors"
	"sponge/internal/dao"
	"sponge/internal/model"
)

type ChatService struct {
	chatDao *dao.ChatDao
	userDao *dao.UserDao
}

func NewChatService(chatDao *dao.ChatDao, userDao *dao.UserDao) *ChatService {
	return &ChatService{
		chatDao: chatDao,
		userDao: userDao,
	}
}

// SendMessage 发送消息
func (s *ChatService) SendMessage(ctx context.Context, fromID, toID int64, msgType int8, content string) (*model.ChatMessage, error) {
	if fromID == toID {
		return nil, errors.New("不能给自己发消息")
	}

	// 检查目标用户是否存在
	_, err := s.userDao.GetUserByID(ctx, toID)
	if err != nil {
		return nil, errors.New("目标用户不存在")
	}

	message := &model.ChatMessage{
		FromID:  fromID,
		ToID:    toID,
		MsgType: msgType,
		Content: content,
		IsRead:  0,
	}

	if err := s.chatDao.CreateMessage(ctx, message); err != nil {
		return nil, err
	}

	// 更新或创建会话（发送方）
	_ = s.chatDao.UpdateChatSession(ctx, fromID, toID, content, false)

	// 更新或创建会话（接收方）
	_ = s.chatDao.UpdateChatSession(ctx, toID, fromID, content, true)

	return message, nil
}

// GetMessageList 获取消息列表
func (s *ChatService) GetMessageList(ctx context.Context, userID, toUserID int64, page, size int) ([]*model.ChatMessage, int64, error) {
	if size <= 0 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}

	messages, total, err := s.chatDao.GetMessageList(ctx, userID, toUserID, page, size)
	if err != nil {
		return nil, 0, err
	}

	// 标记消息为已读（异步）
	go func() {
		_ = s.chatDao.MarkMessagesAsRead(context.Background(), toUserID, userID)
		// 更新会话未读数
		_ = s.chatDao.UpdateChatSession(context.Background(), userID, toUserID, "", false)
	}()

	return messages, total, nil
}

// GetChatSessionList 获取会话列表
func (s *ChatService) GetChatSessionList(ctx context.Context, userID int64, page, size int) ([]*model.ChatSession, int64, error) {
	if size <= 0 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}

	sessions, total, err := s.chatDao.GetChatSessionList(ctx, userID, page, size)
	if err != nil {
		return nil, 0, err
	}

	return sessions, total, nil
}
