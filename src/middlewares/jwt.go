package middlewares

import (
	"fmt"
	"net/http"
	"personal-blog-system/src/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware 验证 Authorization header 并把 userID 放到 context 中：c.Set("userID", uint)
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}
		tokenStr := parts[1]

		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			// 强校验签名算法
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return utils.JWTSecret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		// 解析 user id（存在 Subject）
		uid, err := strconv.ParseUint(claims.Subject, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
			return
		}
		c.Set("userID", uint(uid))
		c.Next()
	}
}
