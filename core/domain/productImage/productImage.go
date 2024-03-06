package productImage

import (
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ID        string
	Productid string
	ImageKey  string
	ImageURL  string
}

func NewProductImage(data ProductImage) (ProductImage, error) {
	return ProductImage{
		ID:        ksuid.New().String(),
		Productid: data.Productid,
		ImageKey:  data.ImageKey,
		ImageURL:  data.ImageURL,
	}, nil
}
