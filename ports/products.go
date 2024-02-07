package ports

import (
	"context"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/product"
)

type ProductRepository interface {
	CountAllProducts(context context.Context) (*int64, error)
	GetAllProducts(context context.Context, params postgres.QueryParams) ([]*product.Product, error)
	GetProductById(ctx context.Context, id string) (*product.Product, error)
	AddProduct(ctx context.Context, p product.Product) error
	EditProduct(ctx context.Context, p product.Product) error
	DeleteProductById(ctx context.Context, p product.Product) error
}
