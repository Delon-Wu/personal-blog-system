package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL      string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string

	Port string

	JwtSecret string
}

var AppConfig Config

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AppConfig.DatabaseHost = os.Getenv("DB_HOST")
	AppConfig.DatabasePort = os.Getenv("DB_PORT")
	AppConfig.DatabaseUser = os.Getenv("DB_USER")
	AppConfig.DatabasePassword = os.Getenv("DB_PASSWORD")
	AppConfig.DatabaseName = os.Getenv("DB_NAME")

	if AppConfig.Port = os.Getenv("PORT"); AppConfig.Port == "" {
		AppConfig.Port = "8080"
	}

	AppConfig.DatabaseURL = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		AppConfig.DatabaseUser,
		AppConfig.DatabasePassword,
		AppConfig.DatabaseHost,
		AppConfig.DatabasePort,
		AppConfig.DatabaseName)

	AppConfig.JwtSecret = os.Getenv("JWT_SECRET")
}
