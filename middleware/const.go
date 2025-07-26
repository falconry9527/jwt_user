package middleware

import "errors"

var SkipAuthPaths = map[string]bool{
	"/api/users/getAll": true,
	"/api/users/login":  true,
}

var (
	ErrInvalidToken = errors.New("invalid token")
	secretKey       = []byte("your-secret-key") // 生产环境应该从配置中获取
)
