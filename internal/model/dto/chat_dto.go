package dto

import "time"

// SendMessageReq 发送消息请求
type SendMessageReq struct {
	ToUserID int64  `json:"to_user_id,string" binding:"required"`
	MsgType  int8   `json:"msg_type" binding:"oneof=1 2"` // 1-文本, 2-图片
	Content  string `json:"content" binding:"required"`
}

// MessageListReq 消息列表请求
type MessageListReq struct {
	ToUserID int64 `form:"to_user_id" json:"to_user_id,string" binding:"required"`
	Page     int   `form:"page" binding:"min=1"`
	Size     int   `form:"size" binding:"min=1,max=100"`
}

// ChatSessionListReq 会话列表请求
type ChatSessionListReq struct {
	Page int `form:"page" binding:"min=1"`
	Size int `form:"size" binding:"min=1,max=100"`
}

// MessageInfoRes 消息信息响应
type MessageInfoRes struct {
	ID        int64     `json:"id,string"`
	FromID    int64     `json:"from_id,string"`
	ToID      int64     `json:"to_id,string"`
	MsgType   int8      `json:"msg_type"`
	Content   string    `json:"content"`
	IsRead    int8      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

// ChatSessionRes 会话信息响应
type ChatSessionRes struct {
	ID          int64        `json:"id,string"`
	PeerID      int64        `json:"peer_id,string"`
	LastMessage string       `json:"last_message"`
	UnreadCount int          `json:"unread_count"`
	UpdatedAt   time.Time    `json:"updated_at"`
	PeerUser    *UserInfoRes `json:"peer_user"` // 对方用户信息
}
