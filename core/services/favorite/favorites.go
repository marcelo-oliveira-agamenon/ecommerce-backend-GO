package favorites

import (
	"context"
	"errors"
	"time"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/favorite"
	"github.com/ecommerce/core/domain/product"
	"github.com/ecommerce/ports"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

var (
	ErrorCreateFavorite = errors.New("adding favorite to user")
)

type Favorite struct {
	ID        string
	UserID    uuid.UUID
	ProductID string
	Product   product.Product
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type API interface {
	AddFavorite(ctx context.Context, f favorite.Favorite) (*favorite.Favorite, error)
	GetFavoriteByUserIdAndProductId(ctx context.Context, userId string, productId string) (*[]favorite.Favorite, error)
	GetFavoriteByUserId(ctx context.Context, userId string, limit int, offset int) (*[]postgres.Favorite, error)
	DeleteFavorite(ctx context.Context, favId string) (bool, error)
}

type FavoriteService struct {
	favoriteRepository ports.FavoriteRepository
}

func NewFavoriteService(fa ports.FavoriteRepository) *FavoriteService {
	return &FavoriteService{
		favoriteRepository: fa,
	}
}
