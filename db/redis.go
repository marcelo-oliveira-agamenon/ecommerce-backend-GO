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
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})

	RedisServer = client
}