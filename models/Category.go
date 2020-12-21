package models

import (
	"time"

	"gorm.io/gorm"
)

// Category struct model
type Category struct {
	gorm.Model
	ID				int
	Name				string
	Image				JSONB				`gorm:"type:jsonb"`
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
}