package product

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyName   = errors.New("empty name")
	ErrorNameTooLong = errors.New("name too long")
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
