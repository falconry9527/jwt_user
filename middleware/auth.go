package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt_user/msg"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 不需要鉴权的url
		if SkipAuthPaths[c.Request.URL.Path] {
			c.Next()
			return
		}
		// 签权的url
		authHeader := c.GetHeader("token")
		if authHeader == "" {
			msg.ErrorCode(c, msg.StatusUnauthorized)
			c.Abort()
			return
		}
		if authHeader == "bbbb" {
			c.Next()
		} else {
			// 鉴权失败
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

// 如果你使用 Gorilla Mux，中间件可以这样写：
/*
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// 将用户信息存入上下文
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
*/
