package user

import (
	"errors"
	"strings"
)

var (
	ErrorHashPassword  = errors.New("creating invalid password")
	ErrorEmptyPassword = errors.New("empty password")
	ErrorMinLength     = errors.New("password must have at least 3 characters")
)

func NewPassword(password string) (string, error) {
	password = strings.TrimSpace(password)
	if len(password) == 0 {
		return "", ErrorEmptyPassword
	}
	if len(password) < 3 {
		return "", ErrorMinLength
	}

	return password, nil
}
