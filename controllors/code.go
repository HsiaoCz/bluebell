package controllors

type ResCode int64

// 定义状态码
const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidAuth
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidAuth:     "无效的Toekn",
}

func (rc ResCode) Msg() string {
	msg, ok := codeMsgMap[rc]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
