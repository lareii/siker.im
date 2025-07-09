package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return env
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
