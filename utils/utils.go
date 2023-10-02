package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file\n")
	}

	return os.Getenv(key)
}