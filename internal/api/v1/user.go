package v1

import (
	"fmt"
	"sponge/internal/model/dto"
	"sponge/internal/service"
	"sponge/pkg/global"
	"sponge/pkg/res" // 假设你的 res 包路径在这里
	"sponge/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
// @Param        data  body      dto.RegisterReq  true  "注册信息"
// @Success      200   {object}  res.Response "注册成功"
// @Failure      400   {object}  res.Response "参数错误"
// @Failure      500   {object}  res.Response "系统异常"
// @Router       /gateway/user/v1/register [post]
func (a *UserApi) Register(c *gin.Context) {
	var req dto.RegisterReq

	// 1. 绑定并校验参数
	if err := c.ShouldBindJSON(&req); err != nil {
		// 调用翻译器
		errMsg := utils.Translate(err)
		global.Logger.Warn("参数校验失败", zap.String("reason", errMsg))
		// 使用统一响应返回翻译后的中文错误
		res.FailMsg(errMsg, c)
		return
	}

	// 2. 调用业务逻辑
	if err := a.userService.Register(c.Request.Context(), req.UserName, req.Password); err != nil {
		res.FailMsg(err.Error(), c)
		return
	}

	// 成功返回
	res.Ok(c)
}

// Login 用户登录
// @Summary      用户登录
// @Description  通过用户名密码登录并返回 JWT
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      dto.LoginReq  true  "登录信息"
// @Success      200   {object}  res.Response "成功"
// @Router       /gateway/user/v1/login [post]
func (a *UserApi) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		// 调用翻译器
		errMsg := utils.Translate(err)
		// 记录带翻译信息的日志，方便后端排查
		global.Logger.Warn("参数错误", zap.String("reason", errMsg))
		// 使用统一响应返回翻译后的中文错误
		res.FailMsg(errMsg, c)
		return
	}

	// 登录
	token, userID, _, err := a.userService.Login(c.Request.Context(), req.UserName, req.Password)
	if err != nil {
		res.FailMsg(err.Error(), c)
		return
	}

	// 查询用户完整信息
	user, err := a.userService.GetUserProfile(c.Request.Context(), userID)
	if err != nil {
		res.FailMsg("获取用户信息失败", c)
		return
	}

	// 使用转换工具将 model 转换为 DTO
	var userInfo dto.UserInfoRes
	if err := utils.ToDTO(user, &userInfo); err != nil {
		res.FailMsg("数据转换失败", c)
		return
	}

	resp := dto.LoginResp{
		Token: token,
		User:  &userInfo,
	}

	res.OkData(resp, c)
}

// UserProfile 获取个人资料
// @Summary      获取当前登录用户信息
// @Description  通过 JWT 解析获取用户 ID，从数据库查询最新资料
// @Tags         用户模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Success      200   {object}  res.Response{data=dto.UserInfoRes} "查询成功"
// @Failure      401   {object}  res.Response "未授权或Token失效"
// @Router       /gateway/user/v1/profile [get]
func (a *UserApi) UserProfile(c *gin.Context) {
	// 从上下文中获取中间件解析出的 UserID
	userID, exists := c.Get("userID")
	fmt.Print("userID:", userID)
	if !exists {
		res.FailMsg("未授权", c)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		res.FailMsg("用户ID格式错误", c)
		return
	}

	// 查询用户信息
	user, err := a.userService.GetUserProfile(c.Request.Context(), userIDStr)
	if err != nil {
		res.FailMsg("获取用户信息失败", c)
		return
	}

	// 使用转换工具将 model 转换为 DTO
	var userInfo dto.UserInfoRes
	if err := utils.ToDTO(user, &userInfo); err != nil {
		res.FailMsg("数据转换失败", c)
		return
	}

	res.OkData(userInfo, c)
}
