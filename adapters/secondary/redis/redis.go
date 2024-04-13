package redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

type redisSession struct {
	UserId    string
	AccessAt  string
	ExpiresAt string
}

func initRedisDatabase() *redis.Client {
	reAddr := os.Getenv("REDIS_ADDR")
	rePass := os.Getenv("REDIS_PASSWORD")

	cli := redis.NewClient(&redis.Options{
		Addr:     reAddr,
		Password: rePass,
		DB:       0,
	})

	return cli
}

func NewRedisRepository() *redis.Client {
	redis := initRedisDatabase()
	return redis
}
