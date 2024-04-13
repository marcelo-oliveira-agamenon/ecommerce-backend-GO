package ports

import (
	"context"
	"time"
)

type RedisService interface {
	StoreUserSession(context context.Context, userId string, expTime time.Time) error
	ValidateSession(context context.Context, userId string) error
}
