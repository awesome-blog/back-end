package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(env string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(env)
}
