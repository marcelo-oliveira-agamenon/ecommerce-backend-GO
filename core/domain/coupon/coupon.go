package coupon

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	ID         string
	Hash       string
	ExpireDate time.Time
	Discount   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func NewCoupon(data Coupon) (Coupon, error) {
	return Coupon{
		ID:         ksuid.New().String(),
		Hash:       data.Hash,
		ExpireDate: data.ExpireDate,
		Discount:   data.Discount,
	}, nil
}
