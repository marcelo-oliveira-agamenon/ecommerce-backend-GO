package db

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

// Redis connection
var (
	RedisServer		*redis.Client
)

// Create new connection for redis server
func RedisConnection() {
	godotenv.Load(".env")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		Password: redisPassword,
		DB: 0,
	})

	RedisServer = client
}