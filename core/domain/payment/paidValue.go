package payment

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrorEmptyPaidValue      = errors.New("empty paid value")
	ErrorConvertingPaidValue = errors.New("converting paid value, check format")
)

func NewPaidValue(paV string) (*float64, error) {
	paV = strings.TrimSpace(paV)
	if len(paV) == 0 {
		return nil, ErrorEmptyPaidValue
	}

	nToVa, err := strconv.ParseFloat(paV, 64)
	if err != nil {
		return nil, ErrorConvertingPaidValue
	}

	return &nToVa, nil
}
