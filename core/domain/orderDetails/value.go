package orderDetails

import (
	"errors"
)

var (
	ErrorValueBelowZero = errors.New("order detail value cannot be negative")
)

func NewValue(toV float64) (*float64, error) {
	if toV < 0 {
		return nil, ErrorValueBelowZero
	}

	return &toV, nil
}
