package controllers

import (
	"personal-blog-system/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

var User = UserController{}

type UserController struct {
	BaseController
}

// UserResponse 是用于 API 返回的 DTO，不包含 Password
type UserResponse struct {
	ID        uint             `json:"id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Posts     []models.Post    `json:"posts"`
	Comments  []models.Comment `json:"comments"`
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		uc.Error(c, err.Error())
		return
	}

	if err := models.CreateUser(&user); err != nil {
		uc.Error(c, "Could not create user")
		return
	}
	// 返回不包含密码的 DTO
	uc.Success(c, "User created")
}

func (uc *UserController) GetUser(c *gin.Context) {
	id, _ := uc.GetIDFromParam(c)
	user, err := models.GetUserByID(id)
	if err != nil {
		uc.Error(c, "User not found")
		return
	}

	// 返回不包含密码的 DTO
	uc.Success(c, UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Posts:     user.Posts,
		Comments:  user.Comments,
	})
}
