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
	ErrorRegisterUserAccess = errors.New("registering user access")
	ErrorGettingKey         = errors.New("getting key info")
	ErrorParsingInfo        = errors.New("incorrect info format")
	ErrorValidatingUser     = errors.New("validating user session")
	ErrorTokenExpiredAt     = errors.New("token expired at: ")
	ErrorDeletingKey        = errors.New("delete key with user id")
	ErrorRegisterHash       = errors.New("save hash access")
	UserSessionKey          = "user-session-"
	ResetPasswordKey        = "reset-hash-"
)

type RedisRepository struct {
	db *redis.Client
}

func NewRedisSessionRepository(dbConn *redis.Client) *RedisRepository {
	return &RedisRepository{
		db: dbConn,
	}
}

func (re *RedisRepository) SaveInfoIntoRedis(context context.Context, key string, data any) error {
	maData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	errM := re.db.Set(context, key, maData, 0).Err()
	if errM != nil {
		return err
	}

	return nil
}

func (re *RedisRepository) RetriveInfoFromRedis(context context.Context, key string, v any) error {
	data, errG := re.db.Get(context, key).Result()
	if errG != nil {
		return ErrorGettingKey
	}

	errU := json.Unmarshal([]byte(data), v)
	if errU != nil {
		return ErrorParsingInfo
	}

	return nil
}

func (re *RedisRepository) StoreUserSession(context context.Context,
	userId string,
	expTime time.Time,
	ip string,
	device string) error {
	ses := redisSession{
		UserId:    userId,
		AccessAt:  time.Now().Round(0).String(),
		ExpiresAt: expTime.Round(0).String(),
		UserIp:    ip,
	}

	key := UserSessionKey + userId
	errM := re.SaveInfoIntoRedis(context, key, ses)
	if errM != nil {
		return ErrorRegisterUserAccess
	}

	return nil
}

func (re *RedisRepository) ValidateSession(context context.Context, userId string) error {
	key := UserSessionKey + userId
	var unMar redisSession
	errG := re.RetriveInfoFromRedis(context, key, unMar)
	if errG != nil {
		return errG
	}

	expAt, errT := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", unMar.ExpiresAt)
	if errT != nil {
		return ErrorValidatingUser
	}
	if time.Now().Round(0).After(expAt) {
		dt := strings.Split(unMar.ExpiresAt, " ")
		return errors.New(ErrorTokenExpiredAt.Error() + dt[0] + " " + dt[1])
	}

	return nil
}

func (re *RedisRepository) ClearUserSession(context context.Context, userId string) error {
	key := UserSessionKey + userId
	_, errD := re.db.Del(context, key).Result()
	if errD != nil {
		return ErrorDeletingKey
	}

	return nil
}

func (re *RedisRepository) SaveResetPasswordInfo(context context.Context, hash string, expiresAt time.Time) error {
	rest := redisResetPass{
		Hash:      hash,
		ExpiresAt: expiresAt.Round(0).String(),
	}

	key := ResetPasswordKey + hash
	errM := re.SaveInfoIntoRedis(context, key, rest)
	if errM != nil {
		return ErrorRegisterHash
	}

	return nil
}

func (re *RedisRepository) ValidateResetPasswordInfo(context context.Context, hash string) error {
	key := ResetPasswordKey + hash
	var unMar redisResetPass
	errG := re.RetriveInfoFromRedis(context, key, unMar)
	if errG != nil {
		return errG
	}

	expAt, errT := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", unMar.ExpiresAt)
	if errT != nil {
		return ErrorValidatingUser
	}
	if time.Now().Round(0).After(expAt) {
		dt := strings.Split(unMar.ExpiresAt, " ")
		return errors.New(ErrorTokenExpiredAt.Error() + dt[0] + " " + dt[1])
	}

	return nil
}
