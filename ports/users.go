package ports

import (
	"context"

	"github.com/ecommerce/core/domain/user"
)

type UserRepository interface {
	AddUser(ctx context.Context, u user.User) error
	FindOneUser(ctx context.Context, email string) (*user.User, error)
}