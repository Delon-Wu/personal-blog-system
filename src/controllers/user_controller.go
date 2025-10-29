package controllers

import (
	"personal-blog-system/src/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
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
	uc.Success(c, user)
}

func (uc *UserController) GetUser(c *gin.Context) {
	id, _ := uc.GetIDFromParam(c)
	user, err := models.GetUserByID(id)
	if err != nil {
		uc.Error(c, "User not found")
	}

	uc.Success(c, user)
}
