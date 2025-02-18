package ports

import (
	"context"
	"time"
)

type RedisService interface {
	StoreUserSession(context context.Context, userId string, expTime time.Time, ip string, device string) error
	ValidateSession(context context.Context, userId string) error
	ClearUserSession(context context.Context, userId string) error
}
