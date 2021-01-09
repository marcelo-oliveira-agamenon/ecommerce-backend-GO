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
	CategoryID			Category				`gorm:"foreignKey:Categoryid;references:ID"`
	Value				float64
	Photos			JSONB					`gorm:"type:jsonb"`
	StockQtd			int
	Description			string
	TypeUnit			string
	TecnicalDetails		string
	HasPromotion		bool
	Discount			float64
	HasShipping			bool
	ShippingPrice		float64
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
	Favorites			[]Favorites				`gorm:"foreignKey:ProductID"`
	Order				[]Order				`gorm:"foreignKey:ProductID"`
}