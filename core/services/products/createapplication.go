package products

import (
	"context"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/product"
)

func (p *ProductService) GetAllProducts(context context.Context,
	limit int,
	offset int,
	getByCategory string,
	getByPromotion string,
	getRecentOnes string,
	getByName string) (*[]product.Product, *int64, error) {
	var params postgres.QueryParamsProduct = postgres.QueryParamsProduct{
		Limit:          limit,
		Offset:         offset,
		GetByCategory:  getByCategory,
		GetByPromotion: getByPromotion,
		GetRecentOnes:  getRecentOnes,
		GetByName:      getByName,
	}

	prod, count, errG := p.productRepository.GetAllProducts(context, params)
	if errG != nil {
		return nil, nil, ErrorGetAllProduct
	}

	return prod, count, nil
}

func (p *ProductService) CreateProduct(context context.Context, data product.Product) (*product.Product, error) {
	pv, errV := ValidateProductFields(data)
	if errV != nil {
		return nil, errV
	}

	errA := p.productRepository.AddProduct(context, *pv)
	if errA != nil {
		return nil, ErrorCreateProduct
	}

	return pv, nil
}

func (p *ProductService) EditProduct(context context.Context, data product.Product) (*product.Product, error) {
	pv, errV := ValidateProductFields(data)
	if errV != nil {
		return nil, errV
	}

	errE := p.productRepository.EditProduct(context, *pv)
	if errE != nil {
		return nil, ErrorEditProduct
	}

	return &data, nil
}

func (p *ProductService) GetProductById(context context.Context, id string) (*product.Product, error) {
	prod, errG := p.productRepository.GetProductById(context, id)
	if errG != nil {
		return nil, ErrorProductDoesntExist
	}

	return prod, nil
}

func (p *ProductService) DeleteProductById(context context.Context, data product.Product) error {
	errD := p.productRepository.DeleteProductById(context, data)
	if errD != nil {
		return ErrorDeleteProduct
	}

	return nil
}
