package orderDetails

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

var (
	ErrorEmptyOrder = errors.New("empty order id")
)

func NewOrderId(orderId string) (uuid.UUID, error) {
	orderId = strings.TrimSpace(orderId)
	if len(orderId) == 0 {
		return uuid.UUID{}, ErrorEmptyOrder
	}

	nOrderId := uuid.FromStringOrNil(orderId)
	return nOrderId, nil
}
