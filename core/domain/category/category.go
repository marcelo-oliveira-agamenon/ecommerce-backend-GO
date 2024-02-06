package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        int
	Name      string
	ImageKey  string
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewCategory(data Category) (Category, error) {
	return Category{
		Name:     data.Name,
		ImageKey: data.ImageKey,
		ImageURL: data.ImageURL,
	}, nil
}
