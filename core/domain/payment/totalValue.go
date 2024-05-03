package payment

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrorEmptyTotalValue      = errors.New("empty total value")
	ErrorConvertingTotalValue = errors.New("converting total value, check format")
	ErrorTotalValueIsZero     = errors.New("total value can't be zero")
)

func NewTotalValue(toVa string) (*float64, error) {
	toVa = strings.TrimSpace(toVa)
	if len(toVa) == 0 {
		return nil, ErrorEmptyTotalValue
	}

	nToVa, err := strconv.ParseFloat(toVa, 64)
	if err != nil {
		return nil, ErrorConvertingTotalValue
	}
	if nToVa == 0 {
		return nil, ErrorTotalValueIsZero
	}

	return &nToVa, nil
}
