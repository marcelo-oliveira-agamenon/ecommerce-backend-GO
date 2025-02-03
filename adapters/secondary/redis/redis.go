package redis

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

type redisSession struct {
	UserId    string
	AccessAt  string
	ExpiresAt string
	UserIp    string
	Device    string
}

func initRedisDatabase() (*redis.Client, error) {
	reAddr := os.Getenv("REDIS_ADDR")
	rePass := os.Getenv("REDIS_PASSWORD")

	cli := redis.NewClient(&redis.Options{
		Addr:     reAddr,
		Password: rePass,
		DB:       0,
	})

	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func NewRedisRepository() (*redis.Client, error) {
	redis, err := initRedisDatabase()
	if err != nil {
		return nil, err
	}
	return redis, nil
}
