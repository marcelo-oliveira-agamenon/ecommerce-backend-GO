package redis

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ErrorUserIdFormat       = errors.New("incorrect user id format")
	ErrorRegisterUserAccess = errors.New("registering user access")
	ErrorGettingKey         = errors.New("validating user session")
	ErrorTokenExpiredAt     = errors.New("token expired at: ")
)

type RedisRepository struct {
	db *redis.Client
}

func NewRedisSessionRepository(dbConn *redis.Client) *RedisRepository {
	return &RedisRepository{
		db: dbConn,
	}
}

func (re *RedisRepository) StoreUserSession(context context.Context, userId string, expTime time.Time) error {
	ses := redisSession{
		UserId:    userId,
		AccessAt:  time.Now().Round(0).String(),
		ExpiresAt: expTime.Round(0).String(),
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

func (re *RedisRepository) ValidateSession(context context.Context, userId string) error {
	data, errG := re.db.Get(context, userId).Result()
	if errG != nil {
		return ErrorGettingKey
	}

	var unMar redisSession
	errU := json.Unmarshal([]byte(data), &unMar)
	if errU != nil {
		return ErrorGettingKey
	}

	expAt, errT := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", unMar.ExpiresAt)
	if errT != nil {
		return ErrorGettingKey
	}
	if time.Now().Round(0).After(expAt) {
		dt := strings.Split(unMar.ExpiresAt, " ")
		return errors.New(ErrorTokenExpiredAt.Error() + dt[0] + " " + dt[1])
	}

	return nil
}
