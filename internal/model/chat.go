package model

import "sponge/pkg/constants"

// ChatSession 会话列表
type ChatSession struct {
	BaseModel
	UserID      int64  `gorm:"uniqueIndex:uk_user_peer;not null" json:"user_id,string"`
	PeerID      int64  `gorm:"uniqueIndex:uk_user_peer;not null" json:"peer_id,string"`
	LastMessage string `gorm:"type:varchar(255)" json:"last_message"`
	UnreadCount int    `gorm:"default:0" json:"unread_count"`
}

// ChatMessage 消息详情
type ChatMessage struct {
	BaseModel
	FromID  int64  `gorm:"index:idx_from_to;index:idx_to_from;not null" json:"from_id,string"`
	ToID    int64  `gorm:"index:idx_from_to;index:idx_to_from;not null" json:"to_id,string"`
	MsgType int8   `gorm:"default:1;comment:1-文本, 2-图片..." json:"msg_type"`
	Content string `gorm:"type:text;not null" json:"content"`
	IsRead  int8   `gorm:"default:0" json:"is_read"`
}

func (ChatSession) TableName() string {
	return constants.ChatSessionTableName
}

func (ChatMessage) TableName() string { return constants.ChatMessageTableName }
