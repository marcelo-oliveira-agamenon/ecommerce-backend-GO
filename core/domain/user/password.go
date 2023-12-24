package user

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorHashPassword = errors.New("creating invalid password")
	ErrorEmptyPassword = errors.New("empty password")
	ErrorMinLength = errors.New("password must have at least 3 characters")
)

func NewPassword(password string) (string, error) {
	password = strings.TrimSpace(password)
	if len(password) == 0 {
		return "", ErrorEmptyPassword
	}
	if len(password) < 3 {
		return "", ErrorMinLength 
	}
// maybe move to another util file
	hashPassword, errPassword := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errPassword != nil {
		return "", ErrorHashPassword
	}

	return string(hashPassword), nil
}