package payment

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	ID            uuid.UUID `gorm:"type:uuid"`
	OrderID       string    `gorm:"column:payment_id"`
	UserID        string    `gorm:"column:user_id"`
	TotalValue    float64
	PaidValue     float64
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

var (
	ErrorUUID           = errors.New("user id")
	ErrorTotalValueLess = errors.New("total value can't be less then paid value")
)

func NewPayment(data Payment) (Payment, error) {
	id, errUUID := uuid.NewV4()
	if errUUID != nil {
		return Payment{}, ErrorUUID
	}

	if data.TotalValue < data.PaidValue {
		return Payment{}, ErrorTotalValueLess
	}

	return Payment{
		ID:            id,
		OrderID:       data.OrderID,
		UserID:        data.UserID,
		TotalValue:    data.TotalValue,
		PaidValue:     data.PaidValue,
		PaymentMethod: data.PaymentMethod,
	}, nil
}
