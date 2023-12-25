package users

import (
	"context"
	"errors"

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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrorMissingFieldsLogin = errors.New("missing fields on login body")
	ErrorUserAlreadyExists  = errors.New("already has a user with this email")
	ErrorUserDoesntExist    = errors.New("user doenst exist")
	ErrorInvalidPassword    = errors.New("invalid password")
)

type API interface {
	Login(context context.Context, body LoginRequest) (*UserResponse, error)
	SignUp(context context.Context, user user.User) (*UserResponse, error)
	DeleteUser(context context.Context, id string) (bool, error)
	UpdateUser(context context.Context, id string, data user.User) (bool, error)
}

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(ur ports.UserRepository) *UserService {
	return &UserService{
		userRepository: ur,
	}
}
