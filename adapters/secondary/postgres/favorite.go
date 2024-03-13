package postgres

import (
	"context"

	"github.com/ecommerce/core/domain/favorite"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// var (
// 	ErrorDeleteFkConstrain = errors.New("cannot delete category with associated products")
// )

type FavoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(dbConn *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{
		db: dbConn,
	}
}

func (fr *FavoriteRepository) AddFavorite(ctx context.Context, f favorite.Favorite) (*favorite.Favorite, error) {
	result := fr.db.Create(&f)
	if result.Error != nil {
		return nil, result.Error
	}

	return &f, nil
}

func (fr *FavoriteRepository) GetFavoriteByUserIdAndProductId(ctx context.Context, prodId string, userId string) (*[]favorite.Favorite, error) {
	var favs []favorite.Favorite

	result := fr.db.Model(favorite.Favorite{
		UserID:    uuid.FromStringOrNil(userId),
		ProductID: prodId,
	}).Find(&favs)
	if result.Error != nil {
		return nil, result.Error
	}

	return &favs, nil
}
