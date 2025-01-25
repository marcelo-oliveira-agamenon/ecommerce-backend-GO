package ordersdetails

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/orderDetails"
	"github.com/ecommerce/ports"
)

var (
	ErrorEmptyOrderDetailsList = errors.New("empty list of orders")
)

type API interface {
	GetTotalOrderValueAndQuatity(ctx context.Context, det []orderDetails.OrderProductData) (*float64, *int, error)
	CheckOrderDetails(ctx context.Context, orderId string, det orderDetails.OrderProductData) (*orderDetails.OrderDetails, error)
}

type OrdersDetailsService struct {
	ordersDetailsRepository ports.OrderDetailsRepository
}

func NewOrdersDetailsService(ord ports.OrderDetailsRepository) *OrdersDetailsService {
	return &OrdersDetailsService{
		ordersDetailsRepository: ord,
	}
}
