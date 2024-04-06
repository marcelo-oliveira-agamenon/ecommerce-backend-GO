package postgres

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/product"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	ErrorProductNotFound = errors.New("product not found")
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(dbConn *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: dbConn,
	}
}

func (pr *ProductRepository) CountAllProducts(context context.Context) (*int64, error) {
	var count int64

	res := pr.db.Table("products").Count(&count)
	if res.Error != nil {
		return nil, res.Error
	}

	return &count, nil
}

func (pr *ProductRepository) GetAllProducts(ctx context.Context, params QueryParamsProduct) (*[]product.Product, *int64, error) {
	var list []product.Product

	query := pr.db.Preload("ProductImage").Joins("Category").Limit(params.Limit).Offset(params.Offset)

	if params.GetByCategory != "" {
		query = query.Where("products.categoryid", params.GetByCategory)
	}

	if params.GetByPromotion != "" {
		query = query.Where("has_promotion", true)
	}

	if params.GetRecentOnes != "" {
		query = query.Order("created_at desc")
	}

	if params.GetByName != "" {
		query = query.Where("products.name like ?", "%"+params.GetByName+"%")
	}

	res := query.Find(&list)
	if res.Error != nil {
		return nil, nil, res.Error
	}

	return &list, &res.RowsAffected, nil
}

func (pr *ProductRepository) GetProductById(ctx context.Context, id string) (*product.Product, error) {
	var prod product.Product
	result := pr.db.Where("products.id", id).First(&prod)
	if result.Error != nil {
		return nil, result.Error
	}

	return &prod, nil
}

func (pr *ProductRepository) AddProduct(ctx context.Context, p product.Product) error {
	result := pr.db.Create(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *ProductRepository) EditProduct(ctx context.Context, p product.Product) error {
	result := pr.db.Save(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *ProductRepository) DeleteProductById(ctx context.Context, p product.Product) error {
	result := pr.db.Unscoped().Delete(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *ProductRepository) CheckProductListById(ctx context.Context, prs pq.StringArray) (*[]string, error) {
	aux := make([]string, len(prs))

	tx := pr.db.Begin()
	for _, s := range prs {
		var p product.Product
		res := tx.Where("id = ?", s).Find(&p)
		if res.Error != nil {
			tx.Rollback()
			return nil, res.Error
		}
		if res.RowsAffected == 0 {
			tx.Rollback()
			return nil, ErrorProductNotFound
		}

		aux = append(aux, p.Name)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &aux, nil
}
