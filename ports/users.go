package ports

import (
	"context"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/user"
)

type UserRepository interface {
	AddUser(ctx context.Context, u user.User) error
	AddBulkUser(ctx context.Context, u []user.User) error
	FindOneUserById(ctx context.Context, id string) (*user.User, error)
	FindOneUserByEmail(ctx context.Context, email string) (*user.User, error)
	FindUsersByFilters(ctx context.Context, params postgres.QueryParamsUsers) (*[]user.User, error)
	DeleteUserById(ctx context.Context, id string) (bool, error)
	UpdateUser(ctx context.Context, id string, u user.User) (bool, error)
}
