package seeds

import (
	"strconv"

	"github.com/ecommerce/core/domain/category"
	"github.com/ecommerce/core/domain/product"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

func SeedInitialData(db *gorm.DB) error {
	// Add category
	cat := []category.Category{
		{
			Name:     "Eletrônicos",
			ImageKey: "test image",
			ImageURL: "test image url",
		},
	}
	catRes := db.Create(&cat)
	if catRes.Error != nil {
		return catRes.Error
	}

	// Add products
	prod := []product.Product{
		{
			ID:              ksuid.New().String(),
			Name:            "Samsung Serie 5 model",
			Categoryid:      strconv.Itoa(cat[0].ID),
			Value:           2000,
			StockQtd:        10,
			HasPromotion:    false,
			Description:     "tv led",
			Discount:        0,
			HasShipping:     false,
			TypeUnit:        "unidade",
			TecnicalDetails: "nice tv led",
			ShippingPrice:   0,
			Rate:            0,
		},
		{
			ID:              ksuid.New().String(),
			Name:            "Samsung Serie 4 model",
			Categoryid:      strconv.Itoa(cat[0].ID),
			Value:           1500,
			StockQtd:        10,
			HasPromotion:    false,
			Description:     "tv led",
			Discount:        0,
			HasShipping:     false,
			TypeUnit:        "unidade",
			TecnicalDetails: "nice tv led",
			ShippingPrice:   0,
			Rate:            0,
		},
	}
	errPr := db.Create(&prod)
	if errPr.Error != nil {
		return errPr.Error
	}

	return nil
}
