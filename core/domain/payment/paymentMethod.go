package payment

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyPaymentMethod   = errors.New("empty payment method")
	ErrorInvalidPaymentMethod = errors.New("invalid payment method")
)

func NewPaymentMethod(paM string) (*string, error) {
	paM = strings.TrimSpace(paM)
	if len(paM) == 0 {
		return nil, ErrorEmptyPaymentMethod
	}
	if paM != "CREDIT_CARD" && paM != "MONEY" && paM != "DEBIT_CARD" && paM != "TRANSFER" && paM != "PIX" {
		return nil, ErrorInvalidPaymentMethod
	}

	return &paM, nil
}
