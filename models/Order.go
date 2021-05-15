package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

//Order struct model
type Order struct {
	gorm.Model
	ID				string
	Userid			uuid.UUID			`gorm:"column:user_id"`
	UserID			User				`gorm:"foreignKey:Userid;references:ID"`
	ProductID			pq.StringArray		`gorm:"type:varchar(64)[]"`
	TotalValue			float64
	Status			string
	Qtd				int
	Paid				bool
	Rate				int
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
	Payment			Payment			`gorm:"foreignKey:OrderID"`
}