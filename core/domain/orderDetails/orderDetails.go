package orderDetails

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type OrderDetails struct {
	gorm.Model
	ID        string
	OrderID   string `gorm:"column:order_id"`
	ProductID string `gorm:"column:product_id"`
	Value     float64
	Qtd       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type OrderProductData struct {
	ProductId  string
	Quantity   int
	Value      float64
	CouponUsed bool
}

func NewOrderDetails(data OrderDetails) (OrderDetails, error) {
	return OrderDetails{
		ID:        ksuid.New().String(),
		ProductID: data.ProductID,
		Value:     data.Value,
		Qtd:       data.Qtd,
		OrderID:   data.OrderID,
	}, nil
}
