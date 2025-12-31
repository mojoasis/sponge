package res

import (
	"net/http"
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
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// Ok 成功返回
func Ok(c *gin.Context) {
	R(xerror.SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkData(data interface{}, c *gin.Context) {
	R(xerror.SUCCESS, data, "查询成功", c)
}

// FailMsg 失败返回

func Fail(c *gin.Context) {
	R(xerror.ERROR, nil, "操作失败", c)
}
func FailMsg(msg string, c *gin.Context) {
	R(xerror.ERROR, map[string]interface{}{}, msg, c)
}

// FailCode 业务错误返回
func FailCode(code int, msg string, c *gin.Context) {
	R(code, map[string]interface{}{}, msg, c)
}
