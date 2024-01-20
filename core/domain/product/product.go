package product

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID              string
	Name            string
	Categoryid      string
	Value           float64
	StockQtd        int
	Description     string
	TypeUnit        string
	TecnicalDetails string
	HasPromotion    bool
	Discount        float64
	HasShipping     bool
	ShippingPrice   float64
	Rate            int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func NewProduct(data Product) (Product, error) {
	return Product{
		ID:              ksuid.New().String(),
		Name:            data.Name,
		Categoryid:      data.Categoryid,
		Value:           data.Value,
		StockQtd:        data.StockQtd,
		Description:     data.Description,
		TypeUnit:        data.TypeUnit,
		TecnicalDetails: data.TecnicalDetails,
		HasPromotion:    data.HasPromotion,
		Discount:        data.Discount,
		HasShipping:     data.HasShipping,
		ShippingPrice:   data.ShippingPrice,
		Rate:            data.Rate,
	}, nil
}
