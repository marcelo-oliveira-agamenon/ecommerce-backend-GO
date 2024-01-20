package ports

import (
	"context"

	"github.com/ecommerce/core/domain/product"
)

type ProductRepository interface {
	AddProduct(ctx context.Context, p product.Product) error
}
