package dto

// LoginReq 登录请求 (全必填)
type LoginReq struct {
	UserName string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required,min=6" example:"123456"`
}

// RegisterReq 注册请求 (部分必填)
type RegisterReq struct {
	UserName string `json:"username" binding:"required,min=4,max=20" example:"newuser"`
	Password string `json:"password" binding:"required,min=6" example:"secret123"`

	// 可选字段：使用指针。如果不传，后端收到的是 nil
	Email  *string `json:"email" binding:"omitempty,email" example:"test@example.com"`
	Avatar *string `json:"avatar" binding:"omitempty,url"`
}

// LoginResp 登录响应
type LoginResp struct {
	Token string       `json:"token"`
	User  *UserInfoRes `json:"user"`
}

// UserInfoRes 用户信息响应
type UserInfoRes struct {
	ID              string `json:"id"`
	UserName        string `json:"user_name"`
	NickName        string `json:"nick_name"`
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
}
