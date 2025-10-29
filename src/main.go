package main

import (
	"os"
	"personal-blog-system/src/config"
	"personal-blog-system/src/database"
	"personal-blog-system/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Connect()
	// 初始化数据库连接（只在入口调用一次）
	database.Connect()
	routes.SetupRoutes(r)

	if port := os.Getenv("PORT"); port != "" {
		r.Run(port)
	} else {
		r.Run("8080")
	}
}
