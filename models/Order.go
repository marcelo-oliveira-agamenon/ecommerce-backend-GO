package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

//Order struct model
type Order struct {
	ID				string
	UserID			uuid.UUID
	ProductID			string
	TotalValue			float64
	Status			string
	Qtd				int
	Paid				bool
	Rate				int
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
}