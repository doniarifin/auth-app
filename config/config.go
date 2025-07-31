package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("fail load .env:", err)
	}
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
