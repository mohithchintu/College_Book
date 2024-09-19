package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI  string
	Port      string
	SecretKey string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig.MongoURI = getEnv("MONGODB_URI", "")
	AppConfig.Port = getEnv("PORT", "8080")
	AppConfig.SecretKey = getEnv("SECRET_KEY", "defaultSecret")

	if AppConfig.MongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
