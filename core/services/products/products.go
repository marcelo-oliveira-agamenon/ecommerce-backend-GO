package products

import (
	"context"
	"errors"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/product"
	"github.com/ecommerce/ports"
)

var (
	ErrorProductDoesntExist = errors.New("product doenst exist with this id")
	ErrorGetAllProduct      = errors.New("fetching products")
	ErrorProductCount       = errors.New("countin products")
	ErrorCreateProduct      = errors.New("insert product")
	ErrorEditProduct        = errors.New("editing product")
	ErrorDeleteProduct      = errors.New("deleting product")
	ErrorCheckProductList   = errors.New("checking product list")
)

type API interface {
	GetAllProducts(context context.Context, limit int, offset int, getByCategory string, getByPromotion string, getRecentOnes string, getByName string) (*[]product.Product, *int64, error)
	GetProductById(context context.Context, id string) (*product.Product, error)
	GetProductCount(context context.Context) (*int64, error)
	GetProductQuantityByCategories(context context.Context) (*[]postgres.CountProducts, *int64, error)
	CreateProduct(context context.Context, data product.Product) (*product.Product, error)
	EditProduct(context context.Context, data product.Product) (*product.Product, error)
	DeleteProductById(context context.Context, data product.Product) error
	CheckProductListById(context context.Context, prList string) (*[]string, error)
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
	name, errN := product.NewName(p.Name)
	if errN != nil {
		return nil, errN
	}

	value, errV := product.NewValue(p.Value)
	if errV != nil {
		return nil, errV
	}

	rate, errR := product.NewRate(p.Rate)
	if errR != nil {
		return nil, errR
	}

	disc, errD := product.NewDiscount(p.Discount)
	if errD != nil {
		return nil, errD
	}

	desc, errDi := product.NewDescription(p.Description)
	if errDi != nil {
		return nil, errDi
	}

	tech, errT := product.NewTechnicalDescription(p.TecnicalDetails)
	if errT != nil {
		return nil, errT
	}

	stock, errS := product.NewStock(p.StockQtd)
	if errS != nil {
		return nil, errS
	}

	p.Name = *name
	p.Value = *value
	p.Discount = *disc
	p.Rate = *rate
	p.Description = *desc
	p.TecnicalDetails = *tech
	p.StockQtd = *stock

	prod, errN := product.NewProduct(p)
	if errN != nil {
		return nil, errN
	}

	return &prod, nil
}
