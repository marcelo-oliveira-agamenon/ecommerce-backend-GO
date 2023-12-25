package user

import "errors"

const (
	masc = "masc"
	fem = "fem"
)

var (
	ErrorInvalidGender = errors.New("gender must be masc or fem value")
)

func NewGender(gender string) (string, error) {
	if validGender := gender == masc || gender == fem; !validGender {
		return "", ErrorInvalidGender
	}

	return gender, nil
}