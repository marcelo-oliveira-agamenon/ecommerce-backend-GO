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

func (cr *CategoryRepository) GetAllCategories(ctx context.Context) (*[]category.Category, error) {
	var list []category.Category
	res := cr.db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}

	return &list, nil
}

func (cr *CategoryRepository) GetCategoryById(ctx context.Context, catId string) (*category.Category, error) {
	var aux category.Category
	result := cr.db.Where("id = ?", catId).First(&aux)
	if result.Error != nil {
		return nil, result.Error
	}

	return &aux, nil
}

func (cr *CategoryRepository) AddCategory(ctx context.Context, c category.Category) (*category.Category, error) {
	result := cr.db.Create(&c)
	if result.Error != nil {
		return nil, result.Error
	}

	return &c, nil
}
