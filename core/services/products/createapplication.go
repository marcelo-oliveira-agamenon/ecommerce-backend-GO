package products

import (
	"context"

	"github.com/ecommerce/core/domain/product"
)

func (p *ProductService) CreateProduct(context context.Context, data product.Product) (*ProductResponse, error) {
	product, errN := product.NewProduct(data)
	if errN != nil {
		return nil, errN
	}

	errA := p.productRepository.AddProduct(context, product)
	if errA != nil {
		return nil, errA
	}

	return NewProductResponse(product), nil
}
