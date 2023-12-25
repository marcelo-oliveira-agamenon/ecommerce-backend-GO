package users

import (
	"context"

	"github.com/ecommerce/core/domain/user"
	"github.com/ecommerce/ports"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

type UserResponse struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Address  string
	Phone    string
	Birthday string
	Gender   string
	Roles    pq.StringArray
}

type API interface {
	SignUp(context context.Context, user user.User) (*UserResponse, error)
	DeleteUser(context context.Context, id string) (bool, error)
	UpdateUser(context context.Context, id string, data user.User) error
}

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(ur ports.UserRepository) *UserService {
	return &UserService{
		userRepository: ur,
	}
}
