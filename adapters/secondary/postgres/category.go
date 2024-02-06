package postgres

import (
	"context"

	"github.com/ecommerce/core/domain/category"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(dbConn *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: dbConn,
	}
}

func (cr *CategoryRepository) GetCategoryById(ctx context.Context, catId string) (*category.Category, error) {
	var aux category.Category
	result := cr.db.Where("id = ?", catId).First(&aux)
	if result.Error != nil {
		return nil, result.Error
	}

	return &aux, nil
}

func (cr *CategoryRepository) AddCategory(ctx context.Context, c category.Category) error {
	result := cr.db.Create(&c)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
