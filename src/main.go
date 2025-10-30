package main

import (
	"personal-blog-system/src/config"
	"personal-blog-system/src/database"
	"personal-blog-system/src/models"
	"personal-blog-system/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Connect()
	// 初始化数据库连接（只在入口调用一次）
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	routes.SetupRoutes(r)

	r.Run(":" + config.AppConfig.Port)
}
