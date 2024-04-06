package orders

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/order"
	"github.com/ecommerce/ports"
)

var (
	ErrorCreateFavorite = errors.New("adding favorite to user")
)

type API interface {
	AddOrder(ctx context.Context, userId string, prodId string, qtd int, toV float64) (*order.Order, error)
	GetByUserId(ctx context.Context, userId string, limit int, offset int) (*[]order.Order, error)
}

type OrderService struct {
	orderRepository ports.OrderRepository
}

func NewOrderService(or ports.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: or,
	}
}
