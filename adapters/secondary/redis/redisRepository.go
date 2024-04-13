package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ErrorUserIdFormat       = errors.New("incorrect user id format")
	ErrorRegisterUserAccess = errors.New("registering user access")
)

type RedisRepository struct {
	db *redis.Client
}

func NewRedisSessionRepository(dbConn *redis.Client) *RedisRepository {
	return &RedisRepository{
		db: dbConn,
	}
}

func (re *RedisRepository) StoreUserSession(context context.Context, userId string, expTime string) error {
	ses := redisSession{
		UserId:    userId,
		AccessAt:  time.Now().String(),
		ExpiresAt: expTime,
	}

	maSes, err := json.Marshal(ses)
	if err != nil {
		return ErrorUserIdFormat
	}

	errM := re.db.Set(context, userId, maSes, 0).Err()
	if errM != nil {
		return ErrorRegisterUserAccess
	}

	return nil
}
