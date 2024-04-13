package ports

import (
	"context"
	"time"

	"github.com/ecommerce/core/domain/order"
)

type OrderRepository interface {
	AddOrder(ctx context.Context, o order.Order) (*order.Order, error)
	GetByUserId(ctx context.Context, userId string, limit int, offset int) (*[]order.Order, error)
	GetById(ctx context.Context, id string) (*order.Order, error)
	GetOrderCount(ctx context.Context) (*int64, error)
	GetPaidOrderCount(ctx context.Context) (*int64, error)
	GetOrdersByPeriod(ctx context.Context, inDate time.Time, enDate time.Time) (*[]order.Order, error)
	UpdateOrderPayment(ctx context.Context, o order.Order, paid bool) (*order.Order, error)
	UpdateOrderRate(ctx context.Context, o order.Order, rate int) (*order.Order, error)
	UpdateOrderStatus(ctx context.Context, o order.Order, status string) (*order.Order, error)
}
