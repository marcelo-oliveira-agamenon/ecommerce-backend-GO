package users

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/user"
)

func (u *UserService) SignUp(context context.Context, data user.User) (*UserResponse, error) {
	alreadyHasUser, errRepo := u.userRepository.FindOneUserByEmail(context, data.Email)
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
	phone, errPhone := user.NewPhone(data.Phone)
	if errPhone != nil {
		return nil, errPhone
	}
	gender, errGen := user.NewGender(data.Gender)
	if errGen != nil {
		return nil, errGen
	}
	data.Password = password
	data.Birthday = birthday
	data.Phone = phone
	data.Gender = gender
	user, errUser := user.NewUser(data)
	if errUser != nil {
		return nil, errUser
	}

	if err := u.userRepository.AddUser(context, user); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		Birthday: user.Birthday,
		Gender:   user.Gender,
		Roles:    user.Roles,
	}, nil
}

func (u *UserService) UpdateUser(context context.Context, id string, data user.User) error {
	u.userRepository.UpdateUser(context, id, data)

	return nil
}

func (u *UserService) DeleteUser(context context.Context, id string) (bool, error) {
	du, err := u.userRepository.DeleteUserById(context, id)
	if err != nil {
		return du, err
	}

	return du, nil
}
