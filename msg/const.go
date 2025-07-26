package msg

const (
	CodeSuccess = 200 // 成功

	CodeError = 400 // 失败
	NotFind   = 401 // 记录没有找到

	CodeInternalError      = 500 // 服务器内部错误
	CodeServiceUnavailable = 503 // 服务不可用
	ParamError             = 504 // 参数错误
	StatusUnauthorized     = 505 // 未登录

)

var codeMsg = map[int]string{
	CodeSuccess:            "Success",
	CodeError:              "error",
	CodeInternalError:      "服务器内部错误",
	CodeServiceUnavailable: "服务不可用",
	ParamError:             "参数错误",
	StatusUnauthorized:     "未登录",
	NotFind:                "记录没有找到",
}

func GetMsg(code int) string {
	msg, ok := codeMsg[code]
	if ok {
		return msg
	}
	return codeMsg[CodeInternalError]
}
