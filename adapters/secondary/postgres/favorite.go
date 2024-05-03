package postgres

import (
	"context"
	"time"

	"github.com/ecommerce/core/domain/favorite"
	"github.com/ecommerce/core/domain/product"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
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

	result := fr.db.Where("product_id", prodId).Find(&favs)
	if result.Error != nil {
		return nil, result.Error
	}

	return &favs, nil
}

func (fr *FavoriteRepository) GetFavoriteByUserId(ctx context.Context, userId string, limit int, offset int) (*[]Favorite, error) {
	var favs []Favorite

	result := fr.db.Preload("Product").Joins("INNER JOIN products ON products.id = favorites.product_id").Where("user_id", userId).Limit(limit).Offset(offset).Find(&favs)
	if result.Error != nil {
		return nil, result.Error
	}

	return &favs, nil
}

func (fr *FavoriteRepository) DeleteFavorite(ctx context.Context, favId string) (bool, error) {
	var fav favorite.Favorite
	result := fr.db.Where("id", favId).Delete(&fav)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}
