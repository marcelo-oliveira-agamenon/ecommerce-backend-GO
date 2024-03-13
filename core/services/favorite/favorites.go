package favorites

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/favorite"
	"github.com/ecommerce/ports"
)

var (
	ErrorCreateFavorite = errors.New("adding favorite to user")
)

type API interface {
	GetFavoriteByUserIdAndProductId(ctx context.Context, userId string, productId string) (*[]favorite.Favorite, error)
	AddFavorite(ctx context.Context, f favorite.Favorite) (*favorite.Favorite, error)
}

type FavoriteService struct {
	favoriteRepository ports.FavoriteRepository
}

func NewFavoriteService(fa ports.FavoriteRepository) *FavoriteService {
	return &FavoriteService{
		favoriteRepository: fa,
	}
}
