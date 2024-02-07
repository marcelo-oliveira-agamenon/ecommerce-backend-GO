package products

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/product"
	"github.com/ecommerce/ports"
)

var (
	ErrorProductDoesntExist = errors.New("product doenst exist with this id")
	ErrorProductCount       = errors.New("product count query")
	ErrorCreateProduct      = errors.New("insert product")
	ErrorEditProduct        = errors.New("editing product")
	ErrorDeleteProduct      = errors.New("deleting product")
)

type API interface {
	GetAllProducts(context context.Context, limit int, offset int, getByCategory string, getByPromotion string, getRecentOnes string, getByName string) ([]*product.Product, *int64, error)
	GetProductById(context context.Context, id string) (*product.Product, error)
	CreateProduct(context context.Context, data product.Product) (*product.Product, error)
	EditProduct(context context.Context, data product.Product) (*product.Product, error)
	DeleteProductById(context context.Context, data product.Product) error
}

type ProductService struct {
	productRepository ports.ProductRepository
}

func NewProductService(pr ports.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: pr,
	}
}

func ValidateProductFields(p product.Product) (*product.Product, error) {
	prod, errN := product.NewProduct(p)
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

	return &product.Product{
		Name:            *name,
		Categoryid:      p.Categoryid,
		Value:           *value,
		StockQtd:        *stock,
		Description:     *desc,
		TypeUnit:        p.TypeUnit,
		TecnicalDetails: *tech,
		HasPromotion:    p.HasPromotion,
		Discount:        *disc,
		HasShipping:     p.HasShipping,
		ShippingPrice:   p.ShippingPrice,
		Rate:            *rate,
	}, nil
}
