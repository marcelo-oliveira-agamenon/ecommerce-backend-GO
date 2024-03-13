package favorite

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	ID        string
	UserID    uuid.UUID
	ProductID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewFavorite(data Favorite) (Favorite, error) {
	return Favorite{
		ID:        ksuid.New().String(),
		UserID:    data.UserID,
		ProductID: data.ProductID,
	}, nil
}
