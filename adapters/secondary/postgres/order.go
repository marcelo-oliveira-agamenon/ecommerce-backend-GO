package postgres

import (
	"context"
	"time"

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

func (or *OrderRepository) GetById(ctx context.Context, id string) (*order.Order, error) {
	od := order.Order{
		ID: id,
	}

	result := or.db.First(&od)
	if result.Error != nil {
		return nil, result.Error
	}

	return &od, nil
}

func (or *OrderRepository) GetOrderCount(ctx context.Context) (*int64, error) {
	var co int64

	res := or.db.Table("orders").Count(&co)
	if res.Error != nil {
		return nil, res.Error
	}

	return &co, nil
}

func (or *OrderRepository) GetPaidOrderCount(ctx context.Context) (*int64, error) {
	var co int64

	res := or.db.Model(&order.Order{}).Where("paid = ?", "true").Count(&co)
	if res.Error != nil {
		return nil, res.Error
	}

	return &co, nil
}

func (or *OrderRepository) GetOrdersByPeriod(ctx context.Context, inDate time.Time, enDate time.Time) (*[]order.Order, error) {
	var ord []order.Order

	res := or.db.Model(&order.Order{}).Where("created_at BETWEEN ? AND ?", inDate, enDate).Find(&ord)
	if res.Error != nil {
		return nil, res.Error
	}

	return &ord, nil
}

func (or *OrderRepository) AddOrder(ctx context.Context, o order.Order) (*order.Order, error) {
	result := or.db.Create(&o)
	if result.Error != nil {
		return nil, result.Error
	}

	return &o, nil
}

func (or *OrderRepository) UpdateOrderPayment(ctx context.Context,
	o order.Order, paid bool) (*order.Order, error) {
	result := or.db.Model(&o).Update("paid", paid)
	if result.Error != nil {
		return nil, result.Error
	}

	return &o, nil
}

func (or *OrderRepository) UpdateOrderRate(ctx context.Context,
	o order.Order, rate int) (*order.Order, error) {
	result := or.db.Model(&o).Update("rate", rate)
	if result.Error != nil {
		return nil, result.Error
	}

	return &o, nil
}

func (or *OrderRepository) UpdateOrderStatus(ctx context.Context,
	o order.Order, status string) (*order.Order, error) {
	result := or.db.Model(&o).Update("status", status)
	if result.Error != nil {
		return nil, result.Error
	}

	return &o, nil
}
