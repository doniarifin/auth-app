package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("fail load .env:", err)
	}
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
