package dao

import (
	"context"
	"errors"
	"sponge/internal/model"
	"sponge/pkg/global"

	"gorm.io/gorm"
)

type ChatDao struct {
	db *gorm.DB
}

func NewChatDao() *ChatDao {
	return &ChatDao{
		db: global.DB,
	}
}

// CreateMessage 创建消息
func (d *ChatDao) CreateMessage(ctx context.Context, message *model.ChatMessage) error {
	return d.db.WithContext(ctx).Create(message).Error
}

// GetMessageList 获取消息列表
func (d *ChatDao) GetMessageList(ctx context.Context, fromID, toID int64, page, size int) ([]*model.ChatMessage, int64, error) {
	var messages []*model.ChatMessage
	var total int64

	// 查询双方的消息（from_id 和 to_id 互换）
	query := d.db.WithContext(ctx).Model(&model.ChatMessage{}).
		Where("(from_id = ? AND to_id = ?) OR (from_id = ? AND to_id = ?)", fromID, toID, toID, fromID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err := query.Order("created_at DESC").Offset(offset).Limit(size).Find(&messages).Error
	return messages, total, err
}

// GetChatSession 获取或创建会话
func (d *ChatDao) GetChatSession(ctx context.Context, userID, peerID int64) (*model.ChatSession, error) {
	var session model.ChatSession
	err := d.db.WithContext(ctx).
		Where("user_id = ? AND peer_id = ?", userID, peerID).
		First(&session).Error

	if err != nil {
		// 如果不存在，创建新会话
		if errors.Is(err, gorm.ErrRecordNotFound) {
			session = model.ChatSession{
				UserID:      userID,
				PeerID:      peerID,
				UnreadCount: 0,
			}
			if err := d.db.WithContext(ctx).Create(&session).Error; err != nil {
				return nil, err
			}
			return &session, nil
		}
		return nil, err
	}
	return &session, nil
}

// UpdateChatSession 更新会话信息
func (d *ChatDao) UpdateChatSession(ctx context.Context, userID, peerID int64, lastMessage string, incrementUnread bool) error {
	updates := map[string]interface{}{
		"last_message": lastMessage,
	}

	if incrementUnread {
		updates["unread_count"] = gorm.Expr("unread_count + 1")
	} else {
		updates["unread_count"] = 0
	}

	return d.db.WithContext(ctx).Model(&model.ChatSession{}).
		Where("user_id = ? AND peer_id = ?", userID, peerID).
		Updates(updates).Error
}

// GetChatSessionList 获取会话列表
func (d *ChatDao) GetChatSessionList(ctx context.Context, userID int64, page, size int) ([]*model.ChatSession, int64, error) {
	var sessions []*model.ChatSession
	var total int64

	query := d.db.WithContext(ctx).Model(&model.ChatSession{}).
		Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err := query.Order("updated_at DESC").Offset(offset).Limit(size).Find(&sessions).Error
	return sessions, total, err
}

// MarkMessagesAsRead 标记消息为已读
func (d *ChatDao) MarkMessagesAsRead(ctx context.Context, fromID, toID int64) error {
	return d.db.WithContext(ctx).Model(&model.ChatMessage{}).
		Where("from_id = ? AND to_id = ? AND is_read = ?", fromID, toID, 0).
		Update("is_read", 1).Error
}
