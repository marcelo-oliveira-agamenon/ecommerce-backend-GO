package ports

import (
	"context"

	"github.com/ecommerce/core/domain/coupon"
)

type CouponRepository interface {
	CreateCoupon(ctx context.Context, c coupon.Coupon) (*coupon.Coupon, error)
	GetCouponByHash(ctx context.Context, hash string) (*[]coupon.Coupon, error)
}
