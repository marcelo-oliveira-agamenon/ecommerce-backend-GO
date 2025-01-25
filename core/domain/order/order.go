package order

import (
	"time"

	"github.com/ecommerce/core/domain/orderDetails"
	"github.com/ecommerce/core/domain/payment"
	"github.com/gofrs/uuid"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           string
	Userid       uuid.UUID `gorm:"column:user_id"`
	TotalValue   float64
	Status       string
	TotalQtd     int
	Paid         bool
	Rate         int
	CouponUsed   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Payment      payment.Payment `gorm:"foreignKey:OrderID"`
	OrderDetails []orderDetails.OrderDetails
}

func NewOrder(data Order) (Order, error) {
	return Order{
		ID:         ksuid.New().String(),
		Userid:     data.Userid,
		TotalValue: data.TotalValue,
		Status:     data.Status,
		TotalQtd:   data.TotalQtd,
		Paid:       data.Paid,
		Rate:       data.Rate,
		CouponUsed: data.CouponUsed,
	}, nil
}
