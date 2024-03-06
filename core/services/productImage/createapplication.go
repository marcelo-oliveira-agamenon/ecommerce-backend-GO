package productImages

import (
	"context"

	"github.com/ecommerce/core/domain/productImage"
)

func (p *ProductImageService) GetProductImageById(ctx context.Context, prodId string) (*productImage.ProductImage, error) {
	prodIm, err := p.productImageRepository.GetProductImageById(ctx, prodId)
	if err != nil {
		return nil, ErrorGettingProductImage
	}

	return prodIm, nil
}

func (p *ProductImageService) CreateProductImage(ctx context.Context, pi productImage.ProductImage) (*productImage.ProductImage, error) {
	imgK, errK := productImage.NewImageKey(pi.ImageKey)
	if errK != nil {
		return nil, errK
	}

	imgU, errU := productImage.NewImageUrl(pi.ImageURL)
	if errU != nil {
		return nil, errU
	}

	prodId, errP := productImage.NewProductId(pi.Productid)
	if errP != nil {
		return nil, errP
	}

	pi.ImageKey = *imgK
	pi.ImageURL = *imgU
	pi.Productid = *prodId
	nPi, errN := productImage.NewProductImage(pi)
	if errN != nil {
		return nil, errN
	}

	_, errA := p.productImageRepository.AddProductImage(ctx, nPi)
	if errA != nil {
		return nil, errA
	}

	return &nPi, nil
}

func (p *ProductImageService) DeleteProductImage(ctx context.Context, prodIm productImage.ProductImage) (bool, error) {
	isDel, err := p.productImageRepository.DeleteProductImage(ctx, prodIm)
	if err != nil {
		return false, ErrorDeleteProductImage
	}

	return isDel, nil
}
