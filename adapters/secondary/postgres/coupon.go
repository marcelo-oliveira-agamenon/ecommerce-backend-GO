package postgres

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/coupon"
	"gorm.io/gorm"
)

var (
	ErrorCreatingCoupon = errors.New("creating coupon")
	ErrorListingCoupon  = errors.New("listing coupons")
)

type CouponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(dbConn *gorm.DB) *CouponRepository {
	return &CouponRepository{
		db: dbConn,
	}
}

func (co *CouponRepository) CreateCoupon(ctx context.Context, c coupon.Coupon) (*coupon.Coupon, error) {
	res := co.db.Create(&c)
	if res.Error != nil {
		return nil, ErrorCreatingCoupon
	}

	return &c, nil
}

func (co *CouponRepository) GetCouponByHash(ctx context.Context, hash string) (*[]coupon.Coupon, error) {
	var coupons []coupon.Coupon

	res := co.db.Where("hash = ?", hash).Find(&coupons)
	if res.Error != nil {
		return nil, ErrorListingCoupon
	}

	return &coupons, nil
}

func (co *CouponRepository) GetOneCouponByHash(ctx context.Context, hash string) (*coupon.Coupon, error) {
	var coupon coupon.Coupon

	res := co.db.Where("hash = ?", hash).First(&coupon)
	if res.Error != nil {
		return nil, ErrorListingCoupon
	}

	return &coupon, nil
}
