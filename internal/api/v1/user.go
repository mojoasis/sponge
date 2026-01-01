package v1

import (
	"fmt"
	"sponge/internal/model/dto"
	"sponge/internal/service"
	"sponge/pkg/global"
	"sponge/pkg/res"
	"sponge/pkg/utils"
	"sponge/pkg/xerror"

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
// @Summary      用户注册
// @Description  根据用户名和密码创建新用户
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      dto.RegisterReq  true  "注册信息"
// @Success      200   {object}  res.Response "注册成功"
// @Failure      400   {object}  res.Response "参数错误"
// @Failure      500   {object}  res.Response "系统异常"
// @Router       /gateway/user/v1/register [post]
func (u *UserApi) Register(c *gin.Context) {
	var req dto.RegisterReq
	// 绑定并校验参数
	if err := c.ShouldBindJSON(&req); err != nil {
		// 调用翻译器
		errMsg := utils.Translate(err)
		global.Logger.Warn("参数校验失败", zap.String("reason", errMsg))
		// 使用统一响应返回翻译后的中文错误
		res.FailMsg(errMsg, c)
		return
	}
	// 注册
	if err := u.userService.Register(c.Request.Context(), &req); err != nil {
		res.FailMsg(err.Error(), c)
		return
	}
	res.Ok(c)
}

// Login 用户登录
// @Summary      用户登录
// @Description  通过用户名密码登录并返回JWT
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
		global.Logger.Error("数据转换失败", zap.Error(err))
		res.FailMsg("数据转换失败", c)
		return
	}
	resp := dto.LoginResp{
		Token: token,
		User:  &userInfo,
	}
	res.OkData(resp, c)
}

// UserProfile 获取个人信息
// @Summary      获取个人信息
// @Description  获取个人信息
// @Tags         用户模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Success      200   {object}  res.Response{data=dto.UserInfoRes} "查询成功"
// @Failure      401   {object}  res.Response "未授权或Token失效"
// @Router       /gateway/user/v1/profile [get]
func (u *UserApi) UserProfile(c *gin.Context) {
	// 从上下文中获取中间件解析出的 UserID
	userID, exists := c.Get("userID")
	if !exists {
		res.FailCode(xerror.ERROR_AUTH_TOKEN, "未授权", c)
		return
	}
	// 查询用户信息
	user, err := u.userService.GetUserProfile(c.Request.Context(), userID.(int64))
	if err != nil {
		res.FailMsg("获取用户信息失败", c)
		return
	}
	// model 转换为 DTO
	var userInfo dto.UserInfoRes
	if err := utils.ToDTO(user, &userInfo); err != nil {
		global.Logger.Error("数据转换失败", zap.Error(err))
		res.FailMsg("数据转换失败", c)
		return
	}
	res.OkData(userInfo, c)
}

// UpdateUserProfile 修改个人信息
// @Summary      修改个人信息
// @Description  修改：昵称、头像、签名、背景图片、邮箱
// @Tags         用户模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        data  body      dto.UpdateUserProfileReq  true  "修改成功"
// @Success      200   {object}  res.Response "修改成功"
// @Router       /gateway/user/v1/update/profile [post]
func (u *UserApi) UpdateUserProfile(c *gin.Context) {
	var req dto.UpdateUserProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailMsg(utils.Translate(err), c)
		return
	}

	// 使用我们之前封装的获取 UID 方式
	userID := res.GetUserID(c)

	// 1. 调用 service，获取更新后的完整实体
	updatedUser, err := u.userService.UpdateUserProfile(c.Request.Context(), userID, &req)
	if err != nil {
		res.FailMsg("更新失败: "+err.Error(), c)
		return
	}

	fmt.Print(updatedUser)

	// 2. 将 model 转换为 DTO 返回给前端
	var userInfo dto.UserInfoRes
	if err := utils.ToDTO(updatedUser, &userInfo); err != nil {
		global.Logger.Error("数据转换失败", zap.Error(err))
		res.FailMsg("数据转换失败", c)
		return
	}

	res.OkDataMsg(userInfo, "修改成功", c)
}

// UpdatePassword 修改密码
// @Summary      修改密码
// @Description  修改密码
// @Tags         用户模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        data  body      dto.UpdatePasswordReq  true  "修改密码信息"
// @Success      200   {object}  res.Response "修改成功"
// @Router       /gateway/user/v1/update/password [post]
func (u *UserApi) UpdatePassword(c *gin.Context) {
	var req dto.UpdatePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		// 调用翻译器
		errMsg := utils.Translate(err)
		// 记录带翻译信息的日志，方便后端排查
		global.Logger.Warn("参数错误", zap.String("reason", errMsg))
		// 使用统一响应返回翻译后的中文错误
		res.FailMsg(errMsg, c)
		return
	}

	// 从 Context 获取当前登录用户 ID (假设在 JWT 中间件已存入)
	uid := c.GetInt64("userID")

	err := u.userService.UpdatePassword(c, uid, &req)
	if err != nil {
		res.FailMsg(err.Error(), c)
		return
	}
	res.Ok(c)
}

// UpdatePhone 修改手机号
// @Summary      修改手机号
// @Description  修改手机号
// @Tags         用户模块
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        data  body      dto.UpdatePhoneReq  true  "修改手机号信息"
// @Success      200   {object}  res.Response "修改成功"
// @Router       /gateway/user/v1/update/phone [post]
func (u *UserApi) UpdatePhone(c *gin.Context) {
	var req dto.UpdatePhoneReq
	if err := c.ShouldBindJSON(&req); err != nil {
		// 调用翻译器
		errMsg := utils.Translate(err)
		// 记录带翻译信息的日志，方便后端排查
		global.Logger.Warn("参数错误", zap.String("reason", errMsg))
		// 使用统一响应返回翻译后的中文错误
		res.FailMsg(errMsg, c)
		return
	}

	uid := c.GetInt64("userID")

	err := u.userService.UpdatePhone(c.Request.Context(), uid, &req)
	if err != nil {
		res.FailMsg(err.Error(), c)
		return
	}
	res.Ok(c)
}
