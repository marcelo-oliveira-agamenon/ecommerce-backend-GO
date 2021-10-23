package models

import (
	"time"

	"gorm.io/gorm"
)

//Product struct model
type Product struct {
	gorm.Model
	ID				string
	Name				string
	Categoryid			string
	Value				float64
	StockQtd			int
	Description			string
	TypeUnit			string
	TecnicalDetails		string
	HasPromotion		bool
	Discount			float64
	HasShipping			bool
	ShippingPrice		float64
	Rate				int
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
	Category			Category				`gorm:"foreignKey:Categoryid"`
	Favorites			[]Favorites				`gorm:"foreignKey:ProductID"`
	ProductImage		[]ProductImage			`gorm:"foreignKey:Productid"`
}