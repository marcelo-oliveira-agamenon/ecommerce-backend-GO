package order

import (
	"errors"
)

var (
	ErrorValueBelowZero = errors.New("order value cannot be negative")
)

func NewTotalValue(toV float64) (*float64, error) {
	if toV < 0 {
		return nil, ErrorValueBelowZero
	}

	return &toV, nil
}
