package coupons

import (
	"context"

	"github.com/ecommerce/core/domain/coupon"
	"github.com/ecommerce/ports"
)

// var (
// 	ErrorCreateFavorite = errors.New("adding favorite to user")
// )

type API interface {
	CreateCoupon(ctx context.Context, hash string, expireDate string, discount string) (*coupon.Coupon, error)
	CheckIfThereIsCouponsByHash(ctx context.Context, hash string) (bool, error)
}

type CouponService struct {
	couponRepository ports.CouponRepository
}

func NewCouponService(co ports.CouponRepository) *CouponService {
	return &CouponService{
		couponRepository: co,
	}
}
