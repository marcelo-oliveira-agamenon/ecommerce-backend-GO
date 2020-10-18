package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Favorites struct model
type Favorites struct {
	gorm.Model
	ID			string
	UserID		uuid.UUID
	ProductID		string
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt
}