package users

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/ecommerce/core/domain/user"
	"github.com/ecommerce/ports"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

var (
	FacebookTokenURL        = "https://graph.facebook.com/me?access_token="
	ErrorUserAlreadyExists  = errors.New("user already exist")
	ErrorUserDoesntExist    = errors.New("user doesnt exist")
	ErrorInvalidPassword    = errors.New("invalid password")
	ErrorUserIsNotAdmin     = errors.New("access denied")
	ErrorInvalidToken       = errors.New("invalid token")
	ErrorPasswordsDontMatch = errors.New("passwords dont match")
	ErrorUpdateUser         = errors.New("updating user")
	ErrorInvalidGender      = errors.New("wrong attribute to gender field")
	ErrorInvalidCsv         = errors.New("generating csv file")
	ErrorReadingUserInfo    = errors.New("reading user information")
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
	IsAdmin  string
}

type LoginFacebook struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type ResetPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Reset    string `json:"reset"`
	Hash     string `json:"hash"`
}

type EmailTemplateResetPassword struct {
	Hash      string
	Name      string
	Email     string
	ExpiredAt time.Time
}

type API interface {
	Login(context context.Context, body LoginRequest) (*UserResponse, error)
	LoginFacebook(context context.Context, body LoginFacebook) (*UserResponse, error)
	ResetPassword(context context.Context, body ResetPassword) (bool, error)
	SignUp(context context.Context, user user.User) (*UserResponse, error)
	DeleteUser(context context.Context, id string) (bool, error)
	UpdateUser(context context.Context, id string, data user.User) (bool, error)
	SendEmailResetPassword(context context.Context, email string) (*EmailTemplateResetPassword, error)
	ToggleRoles(context context.Context, id string) (*user.User, error)
	GetUserById(context context.Context, user_id string) (*user.User, error)
	GetUserByEmail(context context.Context, email string) (*user.User, error)
	GetUserCount(context context.Context) (*int64, error)
	ImportUsers(context context.Context, file multipart.File) (bool, error)
	ExportUsers(context context.Context, createdAtStart string, createdAtEnd string, gender string) ([]byte, error)
}

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(ur ports.UserRepository) *UserService {
	return &UserService{
		userRepository: ur,
	}
}

func NewUserResponse(data user.User) *UserResponse {
	return &UserResponse{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Address:  data.Address,
		Phone:    data.Phone,
		Birthday: data.Birthday,
		Gender:   data.Gender,
		Roles:    data.Roles,
	}
}
