package categories

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/category"
	"github.com/ecommerce/ports"
)

var (
	ErrorGetCategory    = errors.New("fetching category list")
	ErrorGetOneCategory = errors.New("getting one category")
	ErrorCreateCategory = errors.New("adding category")
)

type API interface {
	GetAllCategories(ctx context.Context, limit int, offset int) (*[]category.Category, error)
	GetCategoryById(ctx context.Context, catId string) (*category.Category, error)
	AddCategory(ctx context.Context, c category.Category) (*category.Category, error)
	DeleteCategory(ctx context.Context, c category.Category) (bool, error)
}

type CategoryService struct {
	categoryRepository ports.CategoryRepository
}

func NewCategoryService(cr ports.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: cr,
	}
}
