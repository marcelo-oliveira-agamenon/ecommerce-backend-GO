package product

import (
	"errors"
)

var (
	ErrorIncorrectDiscount = errors.New("incorrect discount")
	ErrorDiscountCantBe100 = errors.New("discount vant be 100%")
)

func NewDiscount(discount float64) (*float64, error) {
	if discount < 0 || discount > 100 {
		return nil, ErrorIncorrectDiscount
	}
	if discount == 100 {
		return nil, ErrorDiscountCantBe100
	}

	return &discount, nil
}
