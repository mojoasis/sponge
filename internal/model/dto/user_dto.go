package dto

// RegisterReq 注册请求
type RegisterReq struct {
	UserName string `json:"username" binding:"required,min=4,max=20" example:"newuser"`
	Password string `json:"password" binding:"required,min=6,max=32" example:"secret123"`
	Phone    string `json:"phone" binding:"required,len=11" example:"18888888888"`
	// Email 注册时可选，但如果传了必须符合格式
	Email *string `json:"email" binding:"omitempty,email" example:"test@example.com"`
}

// LoginReq 登录请求
type LoginReq struct {
	UserName string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required,min=6" example:"123456"`
}

// LoginResp 登录响应
type LoginResp struct {
	Token string       `json:"token"`
	User  *UserInfoRes `json:"user"`
}

// UserInfoRes 用户信息响应 (用于个人信息页、他人主页)
type UserInfoRes struct {
	ID              int64  `json:"id,string"`
	UserName        string `json:"userName"`
	NickName        string `json:"nickName"`
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"backgroundImage"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
}

// UpdateUserProfileReq 更新用户信息请求 (增量更新)
// 全部使用指针类型，以便配合 utils.CopyToModelForUpdate 区分“未传”和“传空”
type UpdateUserProfileReq struct {
	Email           *string `json:"email" binding:"omitempty,email"`
	Avatar          *string `json:"avatar" binding:"omitempty,url"`
	NickName        *string `json:"nickName" binding:"omitempty,max=50"`
	Signature       *string `json:"signature" binding:"omitempty,max=255"`
	BackgroundImage *string `json:"backgroundImage" binding:"omitempty,url"`
}

// UpdatePasswordReq 更新密码请求
type UpdatePasswordReq struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6,max=32,nefield=OldPassword"` // nefield 确保新旧密码不等
}

// UpdatePhoneReq 更新手机号请求
type UpdatePhoneReq struct {
	NewPhone string `json:"newPhone" binding:"required,len=11"`
	Code     string `json:"code" binding:"required,len=6"` // 增加验证码字段，实际业务必需
}
