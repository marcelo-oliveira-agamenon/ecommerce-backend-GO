package ports

import (
	"context"

	"github.com/ecommerce/core/domain/favorite"
)

type FavoriteRepository interface {
	AddFavorite(ctx context.Context, f favorite.Favorite) (*favorite.Favorite, error)
	GetFavoriteByUserIdAndProductId(ctx context.Context, prodId string, userId string) (*[]favorite.Favorite, error)
}
