package coupon

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrorDiscountOutOfRange = errors.New("discount must be between 0 and 100")
	ErrorEmptyDiscount      = errors.New("empty discount")
	ErrorParseDiscount      = errors.New("in discount, check format")
)

func NewDiscount(discount string) (*int, error) {
	discount = strings.TrimSpace(discount)
	if len(discount) == 0 {
		return nil, ErrorEmptyDiscount
	}

	disN, errN := strconv.Atoi(discount)
	if errN != nil {
		return nil, ErrorParseDiscount
	}

	if disN > 100 || disN < 0 {
		return nil, ErrorDiscountOutOfRange
	}

	return &disN, nil
}
