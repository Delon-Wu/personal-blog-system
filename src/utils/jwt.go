package utils

import (
	"personal-blog-system/src/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte(config.AppConfig.JwtSecret) // 应该从环境变量中读取

// GenerateToken - 生成一个带 RegisteredClaims 的 token，subject 使用用户 ID（字符串）
func GenerateToken(userID uint, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatUint(uint64(userID), 10),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}
