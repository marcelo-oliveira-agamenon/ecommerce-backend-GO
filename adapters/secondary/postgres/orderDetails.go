package postgres

import (
	"context"

	"github.com/ecommerce/core/domain/orderDetails"
	"gorm.io/gorm"
)

type OrderDetailsRepository struct {
	db *gorm.DB
}

func NewOrderDetailsRepository(dbConn *gorm.DB) *OrderDetailsRepository {
	return &OrderDetailsRepository{
		db: dbConn,
	}
}

func (od *OrderDetailsRepository) AddOrderDetail(ctx context.Context, det orderDetails.OrderDetails) (*orderDetails.OrderDetails, error) {
	res := od.db.Create(&det)
	if res.Error != nil {
		return nil, res.Error
	}

	return &det, nil
}
