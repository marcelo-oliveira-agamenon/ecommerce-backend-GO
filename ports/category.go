package ports

import (
	"context"

	"github.com/ecommerce/core/domain/category"
)

type CategoryRepository interface {
	GetCategoryById(ctx context.Context, catId string) (*category.Category, error)
	AddCategory(ctx context.Context, c category.Category) error
}
