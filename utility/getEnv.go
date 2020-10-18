package utility

import (
	"os"

	"github.com/joho/godotenv"
)

//GetDotEnv returns a value for a key .env
func GetDotEnv(key string) string {
	godotenv.Load(".env")
	value := os.Getenv(key)

	return value
}