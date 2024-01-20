package postgres

import (
	"context"

	"github.com/ecommerce/core/domain/product"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(dbConn *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: dbConn,
	}
}

func (pr *ProductRepository) AddProduct(ctx context.Context, p product.Product) error {
	result := pr.db.Create(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
