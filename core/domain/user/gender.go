package user

import (
	"errors"
	"strings"
)

const (
	masc = "masc"
	fem  = "fem"
)

var (
	ErrorEmptyGender   = errors.New("empty gender")
	ErrorInvalidGender = errors.New("gender must be masc or fem value")
)

func NewGender(gender string) (string, error) {
	gender = strings.TrimSpace(gender)
	if len(gender) == 0 {
		return "", ErrorEmptyGender
	}
	if validGender := gender == masc || gender == fem; !validGender {
		return "", ErrorInvalidGender
	}

	return gender, nil
}
