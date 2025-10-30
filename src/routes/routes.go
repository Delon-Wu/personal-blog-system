package routes

import (
	"personal-blog-system/src/controllers"
	"personal-blog-system/src/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	RegisterRoutes(r)
}

func RegisterRoutes(r *gin.Engine) {
	authRoot := r.Group("/api", middlewares.JWTAuthMiddleware())
	noAuthRoot := r.Group("/api")

	noAuthRoot.POST("/login", controllers.Auth.Login)

	noAuthRoot.POST("/users", controllers.User.CreateUser)
	authRoot.GET("/users/:id", controllers.User.GetUser)
	authRoot.POST("/post", controllers.Post.Create)
	authRoot.PUT("/post", controllers.Post.Edit)
	authRoot.GET("/post/:id", controllers.Post.Get)
	authRoot.POST("/post/comment/:id", controllers.Post.CreateComment)
	authRoot.GET("/post/comment/:id", controllers.Post.GetComments)
	authRoot.DELETE("/post/comment/:id", controllers.Post.DeleteComment)
	authRoot.GET("/post/list", controllers.Post.List)
	authRoot.DELETE("/post/:id", controllers.Post.Delete)
}
