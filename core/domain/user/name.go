package user

import (
	"errors"
	"strings"
)

var (
	ErrorNameTooLong = errors.New("name too long")
	ErrorEmptyName   = errors.New("empty name")
)

func NewName(name string) (*string, error) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return nil, ErrorEmptyName
	}
	if len(name) > 256 {
		return nil, ErrorNameTooLong
	}

	return &name, nil
}
