package msg

const (
	CodeSuccess            = 200 // 成功
	CodeCreated            = 201 // 创建成功
	CodeAccepted           = 202 // 已接受
	CodeNoContent          = 204 // 无内容
	CodeBadRequest         = 400 // 请求错误
	CodeUnauthorized       = 401 // 未授权
	CodePaymentRequired    = 402 // 需要付款
	CodeForbidden          = 403 // 禁止访问
	CodeNotFound           = 404 // 资源不存在
	CodeMethodNotAllowed   = 405 // 方法不允许
	CodeConflict           = 409 // 冲突
	CodeInternalError      = 500 // 服务器内部错误
	CodeNotImplemented     = 501 // 未实现
	CodeServiceUnavailable = 503 // 服务不可用
	ParamError             = 504 // 参数错误

	CodeError = 600 // 失败
	NotFind   = 601 // 记录没有找到

)

var codeMsg = map[int]string{
	CodeSuccess:            "Success",
	CodeError:              "error",
	CodeAccepted:           "已接受",
	CodeNoContent:          "无内容",
	CodeBadRequest:         "请求参数错误",
	CodeUnauthorized:       "未授权",
	CodePaymentRequired:    "需要付款",
	CodeForbidden:          "禁止访问",
	CodeNotFound:           "资源不存在",
	CodeMethodNotAllowed:   "方法不允许",
	CodeConflict:           "资源冲突",
	CodeInternalError:      "服务器内部错误",
	CodeNotImplemented:     "功能未实现",
	CodeServiceUnavailable: "服务不可用",
	ParamError:             "参数错误",
	NotFind:                "记录没有找到",
	CodeCreated:            "创建成功",
}

func GetMsg(code int) string {
	msg, ok := codeMsg[code]
	if ok {
		return msg
	}
	return codeMsg[CodeInternalError]
}
