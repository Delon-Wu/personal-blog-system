package database

import (
	"fmt"
	"personal-blog-system/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the global DB variable if not already initialized.
// It is safe to call multiple times; subsequent calls will return the existing connection.
func Connect() {
	if DB != nil {
		fmt.Println("已连接")
		return
	}

	dsn := config.AppConfig.DatabaseURL + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 返回错误而不是直接 fatal，这样上层可以决定如何处理
		fmt.Println("failed to connect to database: %w", err)
		return
	}
	fmt.Println("Database connection established")
}
