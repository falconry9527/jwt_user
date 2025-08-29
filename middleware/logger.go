package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RequestLogger 创建一个记录请求信息的中间件
func RequestLoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		startTime := time.Now()

		// 记录请求信息
		requestInfo := map[string]interface{}{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"query":  c.Request.URL.RawQuery,
			"ip":     c.ClientIP(),
			"header": c.Request.Header,
		}

		// 读取请求体（用于记录POST/PUT等请求的参数）
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody)) // 重新放回body
		}

		if len(requestBody) > 0 {
			var bodyMap map[string]interface{}
			if err := json.Unmarshal(requestBody, &bodyMap); err == nil {
				requestInfo["body"] = bodyMap
			} else {
				requestInfo["body"] = string(requestBody)
			}
		}

		// 记录请求开始日志
		logger.WithFields(logrus.Fields{
			"type":    "request_start",
			"request": requestInfo,
		}).Info("Request started")

		// 替换ResponseWriter以捕获响应
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 计算请求处理时间
		duration := time.Since(startTime)

		// 记录响应信息
		responseInfo := map[string]interface{}{
			"status":   c.Writer.Status(),
			"duration": duration.String(),
		}

		// 尝试解析响应体为JSON
		var responseBody interface{}
		if err := json.Unmarshal(blw.body.Bytes(), &responseBody); err == nil {
			responseInfo["body"] = responseBody
		} else {
			responseInfo["body"] = blw.body.String()
		}

		// 记录请求完成日志
		logger.WithFields(logrus.Fields{
			"type":     "request_end",
			"request":  requestInfo,
			"response": responseInfo,
		}).Info("Request completed")
	}
}

// bodyLogWriter 用于捕获响应体
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
