package categories

import (
	"context"

	"github.com/ecommerce/core/domain/category"
)

func (p *CategoryService) GetCategoryById(ctx context.Context, catId string) (*category.Category, error) {
	category, errCa := p.categoryRepository.GetCategoryById(ctx, catId)
	if errCa != nil {
		return nil, errCa
	}

	return category, nil
}

func (p *CategoryService) AddCategory(ctx context.Context, c category.Category) error {
	return nil
}
