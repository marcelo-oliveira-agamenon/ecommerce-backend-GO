package products

import (
	"context"

	"github.com/ecommerce/core/domain/product"
)

func (p *ProductService) CreateProduct(context context.Context, data product.Product) (*ProductResponse, error) {
	prod, errN := product.NewProduct(data)
	if errN != nil {
		return nil, errN
	}

	name, errN := product.NewName(prod.Name)
	if errN != nil {
		return nil, errN
	}

	value, errV := product.NewValue(prod.Value)
	if errV != nil {
		return nil, errV
	}

	rate, errR := product.NewRate(prod.Rate)
	if errR != nil {
		return nil, errR
	}

	disc, errD := product.NewDiscount(prod.Discount)
	if errD != nil {
		return nil, errD
	}

	desc, errDi := product.NewDescription(prod.Description)
	if errDi != nil {
		return nil, errDi
	}

	tech, errT := product.NewTechnicalDescription(prod.TecnicalDetails)
	if errT != nil {
		return nil, errT
	}

	stock, errS := product.NewStock(prod.StockQtd)
	if errS != nil {
		return nil, errS
	}

	prod.Name = *name
	prod.Rate = *rate
	prod.Value = *value
	prod.Discount = *disc
	prod.Description = *desc
	prod.TecnicalDetails = *tech
	prod.StockQtd = *stock

	errA := p.productRepository.AddProduct(context, prod)
	if errA != nil {
		return nil, errA
	}

	return NewProductResponse(prod), nil
}
