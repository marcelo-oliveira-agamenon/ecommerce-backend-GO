package ports

import (
	"context"

	"github.com/ecommerce/core/domain/productImage"
)

type ProductImageRepository interface {
	GetProductImageById(ctx context.Context, prodId string) (*productImage.ProductImage, error)
	AddProductImage(ctx context.Context, pi productImage.ProductImage) (*productImage.ProductImage, error)
	DeleteProductImage(ctx context.Context, prodIm productImage.ProductImage) (bool, error)
}
