package res

import (
	"reflect"
	"sponge/pkg/xerror"

	"github.com/gin-gonic/gin"
)

// Response 标准返回结构
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// R 核心逻辑
func R(code int, data interface{}, msg string, c *gin.Context) {
	// 核心封装：处理 nil 为 {} 或 []
	c.JSON(xerror.SUCCESS, Response{
		Code: code,
		Data: ensureNotNil(data),
		Msg:  msg,
	})
}

// ensureNotNil 确保 data 不为 nil。如果是 nil 则返回空对象 map
func ensureNotNil(data interface{}) interface{} {
	// 1. 直接判断原生 nil
	if data == nil {
		return gin.H{}
	}

	// 2. 使用反射处理指针类型的 nil (比如 var u *UserInfo = nil)
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return gin.H{}
	}

	// 3. 如果是切片且为 nil，建议返回 [] 而不是 {}
	if val.Kind() == reflect.Slice && val.IsNil() {
		return []interface{}{}
	}

	return data
}

// Ok 成功返回 - 默认返回 {}
func Ok(c *gin.Context) {
	R(xerror.SUCCESS, nil, "操作成功", c)
}

// OkMsg 携带自定义消息
func OkMsg(msg string, c *gin.Context) {
	R(xerror.SUCCESS, nil, msg, c)
}

// OkData 成功并携带数据
func OkData(data interface{}, c *gin.Context) {
	R(xerror.SUCCESS, data, "查询成功", c)
}

// OkDataMsg 携带数据并自定义消息
func OkDataMsg(data interface{}, msg string, c *gin.Context) {
	R(xerror.SUCCESS, data, msg, c)
}

// Fail 失败返回
func Fail(c *gin.Context) {
	R(xerror.ERROR, nil, "操作失败", c)
}

// FailMsg 失败并携带自定义消息
func FailMsg(msg string, c *gin.Context) {
	R(xerror.ERROR, nil, msg, c)
}

// FailCode 业务错误返回
func FailCode(code int, msg string, c *gin.Context) {
	R(code, nil, msg, c)
}

// GetUserID 获取当前登录用户的 ID
func GetUserID(c *gin.Context) int64 {
	uid, exists := c.Get("userID")
	if !exists {
		return 0
	}
	id, ok := uid.(int64)
	if !ok {
		return 0
	}
	return id
}
