package controllers

import (
	"errors"
	"personal-blog-system/src/database"
	"personal-blog-system/src/models"
	"personal-blog-system/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Auth 校验
var Auth = AuthController{}

type AuthController struct {
	BaseController
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 示例：校验用户名+密码，返回 JWT
func (a *AuthController) Login(c *gin.Context) {
	db := database.DB
	var req AuthController
	if err := c.ShouldBindJSON(&req); err != nil {
		a.Error(c, err.Error())
		return
	}
	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			a.Error(c, "No this user")
			return
		}
		a.Error(c, "db error")
		return
	}
	// 假设用户密码使用 bcrypt 存储在 user.PasswordHash
	if err := user.CheckPassword(req.Password); err != nil {
		a.Error(c, "Account or Password error")
		return
	}
	// 签发 token（有效期 7 天）
	token, err := utils.GenerateToken(user.ID, 168*time.Hour)
	if err != nil {
		a.Error(c, "failed to generate token")
		return
	}
	a.Success(c, gin.H{"token": token, "expires_in": 7 * 24 * 3600})
}

// GetUserID 从 context 获取 userID 的辅助函数
func (a *AuthController) GetUserID(c *gin.Context) (uint, bool) {
	v, ok := c.Get("userID")
	if !ok {
		return 0, false
	}
	uid, ok := v.(uint)
	return uid, ok
}
