package users

import (
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"math/rand"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/user"
	"github.com/ecommerce/core/util"
	"github.com/lib/pq"
)

func (u *UserService) SignUp(context context.Context, data user.User) (*UserResponse, error) {
	ema, errM := user.NewEmail(data.Email)
	if errM != nil {
		return nil, errM
	}

	eus, errG := u.GetUserByEmail(context, data.Email)
	if eus != nil {
		return nil, ErrorUserAlreadyExists
	}
	if errG != nil && errG != ErrorUserDoesntExist {
		return nil, errG
	}

	nm, errN := user.NewName(data.Name)
	if errN != nil {
		return nil, errN
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

	data.Name = *nm
	data.Password = hashPass
	data.Birthday = birthday
	data.Phone = phone
	data.Gender = gender
	data.Roles = roles
	data.Email = ema
	data.WelcomeEmailSended = false
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

func (u *UserService) GetUserCount(context context.Context) (*int64, error) {
	co, err := u.userRepository.GetUserCount(context)
	if err != nil {
		return nil, err
	}

	return co, nil
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

	us, errG := u.GetUserByEmail(context, body.Email)
	if errG != nil {
		return nil, errG
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

	us, errG := u.GetUserByEmail(context, body.Email)
	if errG != nil {
		return nil, errG
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

	us, errG := u.GetUserByEmail(context, body.Email)
	if errG != nil {
		return false, errG
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
	us, errG := u.GetUserByEmail(context, email)
	if errG != nil {
		return nil, errG
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	hash := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	body := EmailTemplateResetPassword{
		Hash:      hash,
		Name:      us.Name,
		Email:     email,
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

	return &body, nil
}

func (u *UserService) ToggleRoles(context context.Context, id string) (*user.User, error) {
	us, errG := u.GetUserById(context, id)
	if errG != nil {
		return nil, errG
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

func (u *UserService) GetUserById(context context.Context, user_id string) (*user.User, error) {
	us, errRepo := u.userRepository.FindOneUserById(context, user_id)
	if errRepo != nil {
		return nil, errRepo
	}
	if us == nil {
		return nil, ErrorUserDoesntExist
	}

	return us, nil
}

func (u *UserService) GetUserByEmail(context context.Context, email string) (*user.User, error) {
	us, errRepo := u.userRepository.FindOneUserByEmail(context, email)
	if errRepo != nil {
		return nil, errRepo
	}
	if us == nil {
		return nil, ErrorUserDoesntExist
	}

	return us, nil
}

func (u *UserService) ExportUsers(context context.Context,
	createdAtStart string, createdAtEnd string, gender string) ([]byte, error) {
	if gender != "masc" && gender != "fem" && gender != "other" {
		return nil, ErrorInvalidGender
	}

	usr, errU := u.userRepository.FindUsersByFilters(context, postgres.QueryParamsUsers{
		CreatedAtStart: createdAtStart,
		CreatedAtEnd:   createdAtEnd,
		Gender:         gender,
	})
	if errU != nil {
		return nil, errU
	}
	if len(*usr) == 0 {
		return []byte{}, nil
	}

	var keys []string
	byF := new(bytes.Buffer)
	file := csv.NewWriter(byF)

	val := reflect.ValueOf(&user.User{}).Elem()
	for i := 0; i < val.NumField(); i++ {
		keys = append(keys, val.Type().Field(i).Name)
	}

	if keE := file.Write(keys); keE != nil {
		return nil, ErrorInvalidCsv
	}

	for _, v := range *usr {
		var rol string
		for _, v := range v.Roles {
			rol = rol + " " + v
		}

		row := []string{v.ID.String(), v.CreatedAt.String(), v.UpdatedAt.String(), v.DeletedAt.Time.String(), v.Name, v.Email, v.Address, v.ImageKey, v.ImageURL, v.Phone, "", "", v.Birthday, v.Gender, rol}
		if vaEr := file.Write(row); vaEr != nil {
			return nil, ErrorInvalidCsv
		}
	}

	file.Flush()
	if errF := file.Error(); errF != nil {
		return nil, ErrorInvalidCsv
	}

	return byF.Bytes(), nil
}

func (u *UserService) ImportUsers(context context.Context, file multipart.File) (bool, error) {
	var usr []user.User
	scn := bufio.NewScanner(file)

	for scn.Scan() {
		elm := strings.Split(scn.Text(), ",")

		nm, errN := user.NewName(elm[0])
		if errN != nil {
			return false, errN
		}

		ema, errM := user.NewEmail(elm[1])
		if errM != nil {
			return false, errM
		}

		pass, errPass := user.NewPassword(elm[5])
		if errPass != nil {
			return false, errPass
		}

		hashPass, errHs := util.HashPassword(pass)
		if errHs != nil {
			return false, errHs
		}

		bir, errBirth := user.NewBirthday(elm[8])
		if errBirth != nil {
			return false, errBirth
		}

		phone, errPhone := user.NewPhone(elm[5])
		if errPhone != nil {
			return false, errPhone
		}

		gender, errGen := user.NewGender(elm[9])
		if errGen != nil {
			return false, errGen
		}

		roles, errRo := user.NewRoles(strings.Split(elm[10], " "))
		if errRo != nil {
			return false, errRo
		}

		nUs, errU := user.NewUser(user.User{
			Name:       *nm,
			Email:      ema,
			Address:    elm[2],
			ImageKey:   elm[3],
			ImageURL:   elm[4],
			FacebookID: elm[7],
			Password:   hashPass,
			Birthday:   bir,
			Phone:      phone,
			Gender:     gender,
			Roles:      roles,
		})
		if errU != nil {
			return false, errU
		}

		usr = append(usr, nUs)
	}

	if errS := scn.Err(); errS != nil {
		return false, ErrorReadingUserInfo
	}

	err := u.userRepository.AddBulkUser(context, usr)
	if err != nil {
		return false, err
	}

	return true, nil
}
