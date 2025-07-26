package msg

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(200, Response{
		Code: code,
		Data: data,
		Msg:  msg,
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
