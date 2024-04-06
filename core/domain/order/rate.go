package order

import (
	"errors"
)

var (
	ErrorOutOfRangeRate = errors.New("rate must be between 0 and 5")
)

func NewRate(rate int) (*int, error) {
	if rate > 5 && rate < 0 {
		return nil, ErrorEmptyStatus
	}

	return &rate, nil
}
