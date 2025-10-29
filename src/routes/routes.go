package routes

import (
	"personal-blog-system/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", controllers.UserController{}.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
}
