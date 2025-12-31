package v1

// LoginReq 登录请求
type LoginReq struct {
	UserName string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required,min=6" example:"123456"`
}

// RegisterReq 注册请求
type RegisterReq struct {
	UserName string `json:"username" binding:"required,min=4,max=20" example:"newuser"`
	Password string `json:"password" binding:"required,min=6" example:"secret123"`
	// 注册特有字段
	//Email string `json:"email" binding:"required,email" example:"test@example.com"`
}
