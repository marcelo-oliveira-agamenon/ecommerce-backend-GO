package users

import (
	"context"

	"github.com/ecommerce/core/domain/user"
	"github.com/ecommerce/core/util"
)

func (u *UserService) SignUp(context context.Context, data user.User) (*UserResponse, error) {
	alreadyHasUser, errRepo := u.userRepository.FindOneUserByEmail(context, data.Email)
	if errRepo != nil {
		return nil, errRepo
	}
	if alreadyHasUser != nil {
		return nil, ErrorUserAlreadyExists
	}

	password, errPass := user.NewPassword(data.Password)
	if errPass != nil {
		return nil, errPass
	}
	hashPass, errHs := util.HashPassword(password)
	if errHs != nil {
		return nil, errHs
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
	roles, errRo := user.NewRoles(data.Roles)
	if errRo != nil {
		return nil, errRo
	}

	data.Password = hashPass
	data.Birthday = birthday
	data.Phone = phone
	data.Gender = gender
	data.Roles = roles
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

func (u *UserService) UpdateUser(context context.Context, id string, data user.User) (bool, error) {
	uu, err := u.userRepository.UpdateUser(context, id, data)
	if err != nil {
		return uu, err
	}

	return uu, nil
}

func (u *UserService) DeleteUser(context context.Context, id string) (bool, error) {
	du, err := u.userRepository.DeleteUserById(context, id)
	if err != nil {
		return du, err
	}

	return du, nil
}

func (u *UserService) Login(context context.Context, body LoginRequest) (*UserResponse, error) {
	if body.Email == "" || body.Password == "" {
		return nil, ErrorMissingFieldsLogin
	}
	//isAdmin :=
	user, errRepo := u.userRepository.FindOneUserByEmail(context, body.Email)
	if errRepo != nil {
		return nil, errRepo
	}
	if user == nil {
		return nil, ErrorUserDoesntExist
	}

	if err := util.CheckPassword(user.Password, body.Password); err != nil {
		return nil, ErrorInvalidPassword
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
