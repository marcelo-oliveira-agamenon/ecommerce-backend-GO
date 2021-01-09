package models

import (
	"gorm.io/gorm"
)

//ProductImage struct model
type ProductImage struct {
	gorm.Model
	ID				string			
	Productid			string
	ProductID			Product		`gorm:"foreignKey:Productid;references:ID"`
	ImageKey			string
	ImageURL			string
}