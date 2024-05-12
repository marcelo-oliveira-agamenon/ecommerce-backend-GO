package user

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrorEmptyEmail   = errors.New("empty email")
	ErrorInvalidEmail = errors.New("invalid email")
)

func NewEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	if len(email) == 0 {
		return "", ErrorEmptyEmail
	}

	if match, err := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email); err != nil || !match {
		return "", ErrorInvalidEmail
	}

	return email, nil
}
