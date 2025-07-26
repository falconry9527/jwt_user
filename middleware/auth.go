package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwt_user/msg"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不需要鉴权的url
		if SkipAuthPaths[c.Request.URL.Path] {
			c.Next()
			return
		}
		// 模拟一个用户生成jwt,一般在登录的时候生成，这边简化一下
		a, _ := GenerateToken(1, "kekeke")
		fmt.Println(a)

		// 从请求头中获取 token
		tokenString := c.GetHeader("token")
		tokenString = a
		if tokenString == "" {
			msg.ErrorCode(c, msg.StatusUnauthorized)
			c.Abort()
			return
		}
		// 解析和验证token
		claims, err := ParseToken(tokenString)
		if err != nil {
			msg.ErrorCode(c, msg.StatusUnauthorized)
			c.Abort()
			return
		}
		// 把用户的常用信息存入上下文，后面可能会用到
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
