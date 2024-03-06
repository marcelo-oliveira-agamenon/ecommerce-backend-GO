package postgres

import (
	"context"

	"github.com/ecommerce/core/domain/productImage"
	"gorm.io/gorm"
)

type ProductImageRepository struct {
	db *gorm.DB
}

func NewProductImageRepository(dbConn *gorm.DB) *ProductImageRepository {
	return &ProductImageRepository{
		db: dbConn,
	}
}

func (pr *ProductImageRepository) GetProductImageById(ctx context.Context, prodId string) (*productImage.ProductImage, error) {
	var pIm productImage.ProductImage
	res := pr.db.Where("id", prodId).First(&pIm)
	if res.Error != nil {
		return nil, res.Error
	}

	return &pIm, nil
}

func (pr *ProductImageRepository) AddProductImage(context context.Context, pi productImage.ProductImage) (*productImage.ProductImage, error) {
	res := pr.db.Create(&pi)
	if res.Error != nil {
		return nil, res.Error
	}

	return &pi, nil
}

func (pr *ProductImageRepository) DeleteProductImage(ctx context.Context, prodImg productImage.ProductImage) (bool, error) {
	result := pr.db.Unscoped().Delete(&prodImg)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
