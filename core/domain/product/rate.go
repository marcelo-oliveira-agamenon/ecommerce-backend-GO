package product

import (
	"errors"
)

var (
	ErrorRateOutOfRange = errors.New("rate out of range")
)

func NewRate(rate int) (*int, error) {
	if rate < 0 || rate > 5 {
		return nil, ErrorRateOutOfRange
	}

	return &rate, nil
}
