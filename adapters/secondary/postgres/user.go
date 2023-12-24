package postgres

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db	*gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) *UserRepository  {
	return &UserRepository{
		db: dbConn,
	}
}

func (ur *UserRepository) AddUser(ctx context.Context, u user.User) error {
	result := ur.db.Create(&u)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	
	return nil
}

func (ur *UserRepository) FindOneUser(ctx context.Context, email string) (*user.User, error) {
	var aux user.User
	result := ur.db.Where("email = ?", email).First(&aux)
	if result.Error != nil {
		return nil, nil
	}
	
	return &aux, nil
}