package favorites

import (
	"context"

	"github.com/ecommerce/core/domain/favorite"
)

func (f *FavoriteService) GetFavoriteByUserIdAndProductId(ctx context.Context, userId string, productId string) (*[]favorite.Favorite, error) {
	favs, err := f.favoriteRepository.GetFavoriteByUserIdAndProductId(ctx, userId, productId)
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
