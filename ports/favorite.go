package ports

import (
	"context"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/favorite"
)

type FavoriteRepository interface {
	AddFavorite(ctx context.Context, f favorite.Favorite) (*favorite.Favorite, error)
	GetFavoriteByUserIdAndProductId(ctx context.Context, prodId string, userId string) (*[]favorite.Favorite, error)
	GetFavoriteByUserId(ctx context.Context, userId string, limit int, offset int) (*[]postgres.Favorite, error)
	DeleteFavorite(ctx context.Context, favId string) (bool, error)
}
