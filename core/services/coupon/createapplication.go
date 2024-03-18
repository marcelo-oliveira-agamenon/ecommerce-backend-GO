package coupons

import (
	"context"

	"github.com/ecommerce/core/domain/coupon"
)

func (c *CouponService) CheckIfThereIsCouponsByHash(ctx context.Context,
	hash string) (bool, error) {
	coH, err := c.couponRepository.GetCouponByHash(ctx, hash)
	if err != nil {
		return false, err
	}

	return len(*coH) > 0, nil
}

func (c *CouponService) CreateCoupon(ctx context.Context,
	hash string, expireDate string,
	discount string) (*coupon.Coupon, error) {
	h, errH := coupon.NewHash(hash)
	if errH != nil {
		return nil, errH
	}

	t, errT := coupon.NewExpireDate(expireDate)
	if errT != nil {
		return nil, errT
	}

	d, errD := coupon.NewDiscount(discount)
	if errD != nil {
		return nil, errD
	}

	co, errC := coupon.NewCoupon(coupon.Coupon{
		ExpireDate: *t,
		Hash:       *h,
		Discount:   *d,
	})
	if errC != nil {
		return nil, errC
	}

	newC, errN := c.couponRepository.CreateCoupon(ctx, co)
	if errN != nil {
		return nil, errN
	}

	return newC, nil
}
