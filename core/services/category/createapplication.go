package categories

import (
	"context"

	"github.com/ecommerce/core/domain/category"
)

func (p *CategoryService) GetAllCategories(ctx context.Context) (*[]category.Category, error) {
	list, err := p.categoryRepository.GetAllCategories(ctx)
	if err != nil {
		return nil, ErrorGetCategory
	}

	return list, nil
}

func (p *CategoryService) GetCategoryById(ctx context.Context, catId string) (*category.Category, error) {
	category, errCa := p.categoryRepository.GetCategoryById(ctx, catId)
	if errCa != nil {
		return nil, ErrorGetOneCategory
	}

	return category, nil
}

func (p *CategoryService) AddCategory(ctx context.Context, c category.Category) (*category.Category, error) {
	name, errN := category.NewName(c.Name)
	if errN != nil {
		return nil, errN
	}

	imgKey, errK := category.NewImageKey(c.ImageKey)
	if errK != nil {
		return nil, errK
	}

	imgUrl, errU := category.NewImageUrl(c.ImageURL)
	if errU != nil {
		return nil, errU
	}

	c.Name = *name
	c.ImageKey = *imgKey
	c.ImageURL = *imgUrl

	newCat, errC := category.NewCategory(c)
	if errC != nil {
		return nil, errC
	}

	cat, errA := p.categoryRepository.AddCategory(ctx, newCat)
	if errA != nil {
		return nil, ErrorCreateCategory
	}

	return cat, nil
}
