package models

import (
	"time"

	"gorm.io/gorm"
)

// Category struct model
type Category struct {
	ID				int
	Name				string
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
}