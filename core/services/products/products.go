package products

import (
	"context"

	"github.com/ecommerce/core/domain/product"
	"github.com/ecommerce/ports"
)

type ProductResponse struct {
	Name            string
	Categoryid      string
	Value           float64
	StockQtd        int
	Description     string
	TypeUnit        string
	TecnicalDetails string
	HasPromotion    bool
	Discount        float64
	HasShipping     bool
	ShippingPrice   float64
	Rate            int
}

type API interface {
	CreateProduct(context context.Context, data product.Product) (*ProductResponse, error)
}

type ProductService struct {
	productRepository ports.ProductRepository
}

func NewProductService(pr ports.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: pr,
	}
}

func NewProductResponse(data product.Product) *ProductResponse {
	return &ProductResponse{
		Name:            data.Name,
		Categoryid:      data.Categoryid,
		Value:           data.Value,
		StockQtd:        data.StockQtd,
		Description:     data.Description,
		TypeUnit:        data.TypeUnit,
		TecnicalDetails: data.TecnicalDetails,
		HasPromotion:    data.HasPromotion,
		Discount:        data.Discount,
		HasShipping:     data.HasShipping,
		ShippingPrice:   data.ShippingPrice,
		Rate:            data.Rate,
	}
}
