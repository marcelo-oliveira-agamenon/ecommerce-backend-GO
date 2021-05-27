package models

import (
	"time"

	"gorm.io/gorm"
)

// Category struct model
type Coupon struct {
	gorm.Model
	ID				string
	Hash				string
	ValityDate			time.Time
	Discount			int
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
}