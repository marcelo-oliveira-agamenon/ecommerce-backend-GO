package users

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/user"
)

func (u *UserService) SignUp(context context.Context, data user.User) (*UserResponse, error)  {
	alreadyHasUser, errRepo := u.userRepository.FindOneUser(context, data.Email)
	if errRepo != nil {
		return nil, errRepo
	}
	if alreadyHasUser != nil {
		return nil, errors.New("already has a user with this email")
	}

	password, errPass := user.NewPassword(data.Password)
	if errPass != nil {
		return nil, errPass
	}
	birthday, errBirth := user.NewBirthday(data.Birthday)
	if errBirth != nil {
		return nil, errBirth
	}
	data.Password = password
	data.Birthday = birthday
	user, errUser := user.NewUser(data)
	if errUser != nil {
		return nil, errUser
	}	

	if err := u.userRepository.AddUser(context, user); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
		Birthday: user.Birthday,
		Gender: user.Gender,
		Roles: user.Roles,
	}, nil
}