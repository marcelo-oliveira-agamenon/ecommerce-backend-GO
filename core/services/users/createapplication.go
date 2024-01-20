package users

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ecommerce/core/domain/user"
	"github.com/ecommerce/core/util"
	"github.com/lib/pq"
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

	return NewUserResponse(data), nil
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
	_, errEm := user.NewEmail(body.Email)
	if errEm != nil {
		return nil, errEm
	}
	_, errPa := user.NewPassword(body.Password)
	if errPa != nil {
		return nil, errPa
	}

	us, errRepo := u.userRepository.FindOneUserByEmail(context, body.Email)
	if errRepo != nil {
		return nil, errRepo
	}
	if us == nil {
		return nil, ErrorUserDoesntExist
	}

	if body.IsAdmin == "true" && !user.IsRoleAdmin(us.Roles) {
		return nil, ErrorUserIsNotAdmin
	}

	if err := util.CheckPassword(us.Password, body.Password); err != nil {
		return nil, ErrorInvalidPassword
	}

	return NewUserResponse(*us), nil
}

func (u *UserService) LoginFacebook(context context.Context, body LoginFacebook) (*UserResponse, error) {
	_, errEm := user.NewEmail(body.Email)
	if errEm != nil {
		return nil, errEm
	}

	us, errRepo := u.userRepository.FindOneUserByEmail(context, body.Email)
	if errRepo != nil {
		return nil, errRepo
	}
	if us == nil {
		return nil, ErrorUserDoesntExist
	}

	resp, errFa := http.Get(FacebookTokenURL + body.Token)
	if errFa != nil {
		return nil, ErrorInvalidToken
	}
	defer resp.Body.Close()

	return NewUserResponse(*us), nil
}

func (u *UserService) ResetPassword(context context.Context, body ResetPassword) (bool, error) {
	_, errEm := user.NewEmail(body.Email)
	if errEm != nil {
		return false, errEm
	}

	if a := user.ComparePasswords(body.Password, body.Reset); !a {
		return false, ErrorPasswordsDontMatch
	}

	us, errRepo := u.userRepository.FindOneUserByEmail(context, body.Email)
	if errRepo != nil {
		return false, errRepo
	}
	if us == nil {
		return false, ErrorUserDoesntExist
	}

	_, errPa := user.NewPassword(body.Password)
	if errPa != nil {
		return false, errPa
	}
	hashPass, errHs := util.HashPassword(body.Password)
	if errHs != nil {
		return false, errHs
	}

	us.Password = hashPass
	state, errUp := u.UpdateUser(context, us.ID.String(), *us)
	if errUp != nil {
		return false, errUp
	}

	return state, nil
}

func (u *UserService) SendEmailResetPassword(context context.Context, email string) (*EmailTemplateResetPassword, error) {
	us, errRepo := u.userRepository.FindOneUserByEmail(context, email)
	if errRepo != nil {
		return nil, errRepo
	}
	if us == nil {
		return nil, ErrorUserDoesntExist
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	hash := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	body := EmailTemplateResetPassword{
		Hash: hash,
		Name: us.Name,
		Year: strconv.Itoa(time.Now().Year()),
	}

	return &body, nil
}

func (u *UserService) ToggleRoles(context context.Context, id string) (*user.User, error) {
	us, errRepo := u.userRepository.FindOneUserById(context, id)
	if errRepo != nil {
		return nil, errRepo
	}
	if us == nil {
		return nil, ErrorUserDoesntExist
	}

	var roles pq.StringArray
	if len(us.Roles) > 1 {
		roles = pq.StringArray{"user"}
	} else {
		roles = pq.StringArray{"user", "admin"}
	}

	us.Roles = roles
	_, err := u.userRepository.UpdateUser(context, id, *us)
	if err != nil {
		return nil, ErrorUpdateUser
	}

	return us, nil
}
