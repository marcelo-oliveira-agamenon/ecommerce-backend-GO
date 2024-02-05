package product

import (
	"errors"
)

var (
	ErrorValueOutOfRange = errors.New("value out of range")
)

func NewValue(value float64) (*float64, error) {
	if value > 10000000 {
		return nil, ErrorValueOutOfRange
	}

	return &value, nil
}
