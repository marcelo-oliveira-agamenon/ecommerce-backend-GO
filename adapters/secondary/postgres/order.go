package postgres

import (
	"context"

	"github.com/ecommerce/core/domain/order"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(dbConn *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: dbConn,
	}
}

func (or *OrderRepository) GetByUserId(ctx context.Context,
	userId string, limit int, offset int) (*[]order.Order, error) {
	var ods []order.Order

	result := or.db.Where("user_id", userId).Limit(limit).Offset(offset).Order("created_at desc").Find(&ods)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ods, nil
}

func (or *OrderRepository) AddOrder(ctx context.Context, o order.Order) (*order.Order, error) {
	result := or.db.Create(&o)
	if result.Error != nil {
		return nil, result.Error
	}

	return &o, nil
}
