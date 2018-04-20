package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvStr(key, val string) string {
	if val != "" {
		return val
	}
	return os.Getenv(key)
}
