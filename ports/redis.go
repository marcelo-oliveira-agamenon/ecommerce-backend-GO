package ports

import (
	"context"
)

type RedisService interface {
	StoreUserSession(context context.Context, userId string, expTime string) error
}
