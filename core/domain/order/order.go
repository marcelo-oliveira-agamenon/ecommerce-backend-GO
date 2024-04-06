package order

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         string
	Userid     uuid.UUID      `gorm:"column:user_id"`
	ProductID  pq.StringArray `gorm:"type:varchar(64)[]"`
	TotalValue float64
	Status     string
	Qtd        int
	Paid       bool
	Rate       int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	// Payment    Payment `gorm:"foreignKey:OrderID"`
}

func NewOrder(data Order) (Order, error) {
	return Order{
		ID:         ksuid.New().String(),
		ProductID:  data.ProductID,
		Userid:     data.Userid,
		TotalValue: data.TotalValue,
		Status:     data.Status,
		Qtd:        data.Qtd,
		Paid:       data.Paid,
		Rate:       data.Rate,
	}, nil
}
