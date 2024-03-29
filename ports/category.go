package ports

import (
	"context"

	"github.com/ecommerce/core/domain/category"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context, limit int, offset int) (*[]category.Category, error)
	GetCategoryById(ctx context.Context, catId string) (*category.Category, error)
	AddCategory(ctx context.Context, c category.Category) (*category.Category, error)
	DeleteCategory(ctx context.Context, c category.Category) (bool, error)
}
