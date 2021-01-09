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
	ImageKey			string
	ImageURL			string
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
}