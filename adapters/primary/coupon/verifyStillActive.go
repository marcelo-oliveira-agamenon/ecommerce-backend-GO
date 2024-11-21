package coupons

import (
	"errors"

	coupons "github.com/ecommerce/core/services/coupon"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingHashParams = errors.New("missing hash in query parameter")
)

func VerifyCouponStillActive(couponAPI coupons.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		hash := ctx.Query("hash")
		if hash == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingExpireDate.Error(),
			})
		}

		co, isValid, err := couponAPI.VerifyIfCouponIsActive(ctx.Context(), hash)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(200).JSON(&fiber.Map{
			"valid":    isValid,
			"discount": co.Discount,
		})
	}
}
