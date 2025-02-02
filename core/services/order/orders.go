package orders

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/order"
	"github.com/ecommerce/core/domain/orderDetails"
	"github.com/ecommerce/ports"
)

var (
	ErrorCreateFavorite = errors.New("adding favorite to user")
)

type API interface {
	AddOrder(ctx context.Context, userId string, qtd int, toV float64, det []orderDetails.OrderDetails, isCouponUsed bool) (*order.Order, error)
	GetById(ctx context.Context, id string) (*order.Order, error)
	GetByUserId(ctx context.Context, userId string, limit int, offset int) (*[]order.Order, error)
	GetOrderCount(ctx context.Context) (*int64, *int64, error)
	GetOrdersByPeriod(ctx context.Context) (*[]OrderMonthQuantity, error)
	GetProfitByOrdersByMonths(ctx context.Context) (*[]MonthData, error)
	UpdatePayment(ctx context.Context, id string, paid string) (*order.Order, error)
	UpdateRate(ctx context.Context, id string, rate string) (*order.Order, error)
	UpdateStatus(ctx context.Context, id string, status string) (*order.Order, error)
}

type OrderService struct {
	orderRepository ports.OrderRepository
}

type OrderMonthQuantity struct {
	Month    string
	Quantity int64
}

type OrderTotalMonth struct {
	OrderId  string
	Subtotal float64
	Month    string
}

type MonthData struct {
	Month string
	Data  []OrderTotalMonth
}

func NewOrderService(or ports.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: or,
	}
}
