package msg

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Response struct {
	TraceId string      `json:"trace_id"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

func Result(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(200, Response{
		TraceId: GetTraceId(c.Request.Context()),
		Code:    code,
		Data:    data,
		Msg:     msg,
	})
}

func Success(c *gin.Context, data interface{}) {
	Result(c, CodeSuccess, data, "SUCCESS")
}

// 内部服务的错误，需要封装给前段
func Error(c *gin.Context, err error) {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorCode(c, NotFind)
		}
		return
	}
	Result(c, CodeError, nil, err.Error())
}

// 业务逻辑上的
func ErrorCode(c *gin.Context, code int) {
	Result(c, code, nil, GetMsg(code))
}

// GetTraceId 获取链路追踪id
func GetTraceId(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}
	return ""
}
