package products

import (
	"context"
	"strings"

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

func (p *ProductService) GetProductCount(context context.Context) (*int64, error) {
	co, err := p.productRepository.CountAllProducts(context)
	if err != nil {
		return nil, err
	}

	return co, nil
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

func (p *ProductService) CheckProductListById(context context.Context, prList string) (*[]string, error) {
	prod := strings.Split(prList, ",")

	li, err := p.productRepository.CheckProductListById(context, prod)
	if err != nil {
		return nil, ErrorCheckProductList
	}

	return li, nil
}

func (p *ProductService) GetProductQuantityByCategories(context context.Context) (*[]postgres.CountProducts, *int64, error) {
	var tot int64

	co, err := p.productRepository.CountProductsByCategories(context)
	if err != nil {
		return nil, nil, err
	}
	if len(*co) == 0 {
		return co, &tot, nil
	}

	aux := *co
	for i := 0; i < len(aux); i++ {
		tot = tot + aux[i].Count
	}

	return co, &tot, nil
}

func (p *ProductService) GetProductsByCategory(context context.Context, categoryId string, limit int,
	offset int) (*[]product.Product, error) {
	var params postgres.QueryParams = postgres.QueryParams{
		Limit:  limit,
		Offset: offset,
	}

	pc, err := p.productRepository.GetProductsByCategory(context, categoryId, params)
	if err != nil {
		return nil, err
	}

	return pc, nil
}
