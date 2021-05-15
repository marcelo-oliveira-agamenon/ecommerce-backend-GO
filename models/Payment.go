package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

//Payment struct model
type Payment struct {
	gorm.Model

	ID				uuid.UUID		`gorm:"type:uuid"`
	OrderID			string
	UserID			string
	TotalValue			float64
	PaidValue			float64
	PaymentMethod		string
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
}