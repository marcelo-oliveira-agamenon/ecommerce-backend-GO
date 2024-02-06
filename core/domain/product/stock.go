package product

import (
	"errors"
)

var (
	ErrorStockOutOfRange = errors.New("stock quantity out of range")
)

func NewStock(stock int) (*int, error) {
	if stock < 0 || stock > 100000 {
		return nil, ErrorEmptyDescription
	}

	return &stock, nil
}
