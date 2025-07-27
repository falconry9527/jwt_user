package jwt

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")
	secretKey       = []byte("your-secret-key") // 生产环境应该从配置中获取
)
