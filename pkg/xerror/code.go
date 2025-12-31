package xerror

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	// 10000+ 为用户模块错误
	ERROR_USER_EXIST     = 10001
	ERROR_USER_NOT_EXIST = 10002
	ERROR_PASSWORD_WRONG = 10003
	ERROR_AUTH_TOKEN     = 10004
	ERROR_AUTH_TIMEOUT   = 10005
)

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "请求参数错误",
	ERROR_USER_EXIST:     "用户已存在",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_AUTH_TOKEN:     "Token鉴权失败",
	ERROR_AUTH_TIMEOUT:   "Token已过期",
}

// GetMsg 获取错误描述
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
