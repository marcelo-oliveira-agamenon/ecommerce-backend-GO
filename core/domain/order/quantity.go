package order

import (
	"errors"
)

var (
	ErrorQuantityAboveLimit = errors.New("quantity of products cannot be above 9999")
	ErrorQuantityBelowZero  = errors.New("quantity of products cannot be negative")
)

func NewQuantity(qtd int) (*int, error) {
	if qtd < 0 {
		return nil, ErrorQuantityBelowZero
	}

	if qtd > 9999 {
		return nil, ErrorQuantityAboveLimit
	}

	return &qtd, nil
}
