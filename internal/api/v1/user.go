package v1

import (
	"sponge/internal/service"
	"sponge/pkg/res" // 假设你的 res 包路径在这里

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	userService *service.UserService // 明确依赖 Service
}

// NewUserApi 构造函数，Wire 会自动注入 service
func NewUserApi(svc *service.UserService) *UserApi {
	return &UserApi{userService: svc}
}

// Register 用户注册
// @Summary      用户注册接口
// @Description  根据用户名和密码创建新用户，ID 使用 UUID V5 自动生成
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      RegisterReq  true  "注册信息"
// @Success      200   {object}  res.Response "注册成功"
// @Failure      400   {object}  res.Response "参数错误"
// @Failure      500   {object}  res.Response "系统异常"
// @Router       /gateway/user/v1/register [post]
func (a *UserApi) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailMsg("参数错误", c)
		return
	}
	if err := a.userService.Register(req.UserName, req.Password); err != nil {
		res.FailMsg(err.Error(), c)
		return
	}
	res.Ok(c)
}

// Login 用户登录
// @Summary      用户登录
// @Description  通过用户名密码登录并返回 JWT
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      LoginReq  true  "登录信息"
// @Success      200   {object}  res.Response "成功"
// @Router       /gateway/user/v1/login [post]
func (a *UserApi) Login(c *gin.Context) {
	var req LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailMsg("参数错误", c)
		return
	}

	// 登录
	token, user, err := a.userService.Login(c, req.UserName, req.Password)
	if err != nil {
		res.FailMsg(err.Error(), c)
		return
	}
	res.OkData(gin.H{"token": token, "user": user}, c)
}

// UserProfile 获取个人资料
// @Summary      获取当前登录用户信息
// @Description  通过 JWT 解析获取用户 ID，从数据库查询最新资料
// @Tags         用户模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Success      200   {object}  res.Response{data=model.User} "查询成功"
// @Failure      401   {object}  res.Response "未授权或Token失效"
// @Router       /gateway/user/v1/profile [get]
func (a *UserApi) UserProfile(c *gin.Context) {
	// 从上下文中获取中间件解析出的 UserID
	userID, _ := c.Get("userID")
	// 简单演示，实际应该去 DB 查最新信息
	res.OkData(gin.H{"user_id": userID, "status": "authenticated"}, c)
}
