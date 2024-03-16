package favorites

import (
	"context"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/favorite"
)

func (f *FavoriteService) GetFavoriteByUserIdAndProductId(ctx context.Context, userId string, productId string) (*[]favorite.Favorite, error) {
	favs, err := f.favoriteRepository.GetFavoriteByUserIdAndProductId(ctx, productId, userId)
	if err != nil {
		return nil, err
	}

	return favs, nil
}

func (f *FavoriteService) GetFavoriteByUserId(ctx context.Context, userId string, limit int, offset int) (*[]postgres.Favorite, error) {
	favs, err := f.favoriteRepository.GetFavoriteByUserId(ctx, userId, limit, offset)
	if err != nil {
		return nil, err
	}

	return favs, nil
}

func (f *FavoriteService) AddFavorite(ctx context.Context, fa favorite.Favorite) (*favorite.Favorite, error) {
	uId, errU := favorite.NewUserId(fa.UserID)
	if errU != nil {
		return nil, errU
	}

	pId, errP := favorite.NewProductId(fa.ProductID)
	if errP != nil {
		return nil, errP
	}

	fa.UserID = *uId
	fa.ProductID = *pId

	nFav, errN := favorite.NewFavorite(fa)
	if errN != nil {
		return nil, errN
	}

	fav, errF := f.favoriteRepository.AddFavorite(ctx, nFav)
	if errF != nil {
		return nil, ErrorCreateFavorite
	}

	return fav, nil
}

func (f *FavoriteService) DeleteFavorite(ctx context.Context, favId string) (bool, error) {
	isDel, err := f.favoriteRepository.DeleteFavorite(ctx, favId)
	if err != nil {
		return false, err
	}

	return isDel, nil
}
