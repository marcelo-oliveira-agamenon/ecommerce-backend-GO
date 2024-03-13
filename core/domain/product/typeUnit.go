package product

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyType   = errors.New("empty type unit")
	ErrorTypeTooLong = errors.New("type unit too long")
)

func NewTypeUnit(typeUnit string) (*string, error) {
	typeUnit = strings.TrimSpace(typeUnit)
	if len(typeUnit) == 0 {
		return nil, ErrorEmptyType
	}
	if len(typeUnit) > 256 {
		return nil, ErrorTypeTooLong
	}

	return &typeUnit, nil
}
