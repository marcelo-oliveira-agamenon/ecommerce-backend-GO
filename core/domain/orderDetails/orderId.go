package orderDetails

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyOrder = errors.New("empty order id")
)

func NewOrderId(orderId string) (string, error) {
	orderId = strings.TrimSpace(orderId)
	if len(orderId) == 0 {
		return "", ErrorEmptyOrder
	}

	return orderId, nil
}
