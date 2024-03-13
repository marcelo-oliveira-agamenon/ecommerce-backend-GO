package productImages

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/productImage"
	"github.com/ecommerce/ports"
)

var (
	ErrorGettingProductImage = errors.New("failed getting product image with this id")
	ErrorDeleteProductImage  = errors.New("failed deleting product image with this id")
)

type API interface {
	GetProductImageById(ctx context.Context, prodId string) (*productImage.ProductImage, error)
	CreateProductImage(ctx context.Context, pi productImage.ProductImage) (*productImage.ProductImage, error)
	DeleteProductImage(ctx context.Context, prodIm productImage.ProductImage) (bool, error)
}

type ProductImageService struct {
	productImageRepository ports.ProductImageRepository
}

func NewProductImageService(pir ports.ProductImageRepository) *ProductImageService {
	return &ProductImageService{
		productImageRepository: pir,
	}
}
