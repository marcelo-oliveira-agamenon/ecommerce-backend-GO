package postgres

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/user"
	"gorm.io/gorm"
)

var (
	ErrorUsersFilterNotFound = errors.New("no user found with this filter")
	ErrorUpdatingUserEmail   = errors.New("updating user email status")
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) *UserRepository {
	return &UserRepository{
		db: dbConn,
	}
}

func (ur *UserRepository) AddUser(ctx context.Context, u user.User) error {
	result := ur.db.Create(&u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) AddBulkUser(ctx context.Context, u []user.User) error {
	res := ur.db.Create(&u)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (ur *UserRepository) GetUserCount(ctx context.Context) (*int64, error) {
	var co int64
	res := ur.db.Table("users").Count(&co)
	if res.Error != nil {
		return nil, res.Error
	}

	return &co, nil
}

func (ur *UserRepository) FindOneUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var aux user.User
	result := ur.db.Where("email = ?", email).First(&aux)
	if aux.Email == "" {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &aux, nil
}

func (ur *UserRepository) DeleteUserById(ctx context.Context, id string) (bool, error) {
	result := ur.db.Delete(&user.User{}, &id)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, id string, u user.User) (bool, error) {
	result := ur.db.Where(&id).Updates(&u)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func (ur *UserRepository) FindOneUserById(ctx context.Context, id string) (*user.User, error) {
	var user user.User
	result := ur.db.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) FindUsersByFilters(ctx context.Context, params QueryParamsUsers) (*[]user.User, error) {
	var usr []user.User
	baQu := ur.db

	if params.CreatedAtStart != "" && params.CreatedAtEnd != "" {
		baQu = baQu.Where("created_at BETWEEN ?::date AND ?::date", params.CreatedAtStart, params.CreatedAtEnd)
	}

	if params.Gender != "" {
		baQu = baQu.Where("gender = ?", params.Gender)
	}

	res := baQu.Find(&usr)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, ErrorUsersFilterNotFound
	}

	return &usr, nil
}

func (ur *UserRepository) WelcomeEmailReceived(message string) error {
	var us user.User
	us.WelcomeEmailSended = true
	usEmail := message

	err := ur.db.Where("email = ?", usEmail).Updates(&us)
	if err.Error != nil {
		return ErrorUpdatingUserEmail
	}

	return nil
}
