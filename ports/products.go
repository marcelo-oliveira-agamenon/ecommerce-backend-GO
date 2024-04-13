package ports

import (
	"context"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/product"
	"github.com/lib/pq"
)

type ProductRepository interface {
	CountAllProducts(context context.Context) (*int64, error)
	CountProductsByCategories(context context.Context) (*[]postgres.CountProducts, error)
	GetAllProducts(context context.Context, params postgres.QueryParamsProduct) (*[]product.Product, *int64, error)
	GetProductById(ctx context.Context, id string) (*product.Product, error)
	AddProduct(ctx context.Context, p product.Product) error
	EditProduct(ctx context.Context, p product.Product) error
	DeleteProductById(ctx context.Context, p product.Product) error
	CheckProductListById(ctx context.Context, prs pq.StringArray) (*[]string, error)
}
