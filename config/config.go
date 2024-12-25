package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort       string
	DatabasePath  string
	JWTSecret     string
	Environment   string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	AppConfig = Config{
		AppPort:       getEnv("APP_PORT", "8080"),
		DatabasePath:  getEnv("DATABASE_PATH", "healthcare.db"),
		JWTSecret:     getEnv("JWT_SECRET", "your_jwt_secret"),
		Environment:   getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
