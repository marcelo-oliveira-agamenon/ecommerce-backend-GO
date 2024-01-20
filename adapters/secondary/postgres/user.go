package postgres

import (
	"context"

	"github.com/ecommerce/core/domain/user"
	"gorm.io/gorm"
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
