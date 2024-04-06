package ports

import (
	"context"

	"github.com/ecommerce/core/domain/order"
)

type OrderRepository interface {
	AddOrder(ctx context.Context, o order.Order) (*order.Order, error)
	GetByUserId(ctx context.Context, userId string, limit int, offset int) (*[]order.Order, error)
}
