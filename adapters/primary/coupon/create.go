package coupons

import (
	"errors"

	coupons "github.com/ecommerce/core/services/coupon"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingExpireDate  = errors.New("missing expire date")
	ErrorMissingDiscount    = errors.New("missing discount")
	ErrorMissingTitle       = errors.New("title")
	ErrorCouponAlreadyExist = errors.New("coupon already exist with this title")
)

func CreateCoupon(couponAPI coupons.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		exp := ctx.FormValue("expire_date")
		if exp == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingExpireDate.Error(),
			})
			return
		}

		dis := ctx.FormValue("discount")
		if dis == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingDiscount.Error(),
			})
			return
		}

		tit := ctx.FormValue("title")
		if tit == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingTitle.Error(),
			})
			return
		}

		hasCoupon, errH := couponAPI.CheckIfThereIsCouponsByHash(ctx.Context(), tit)
		if errH != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errH.Error(),
			})
			return
		}
		if hasCoupon {
			ctx.Status(409).JSON(&fiber.Map{
				"error": ErrorCouponAlreadyExist.Error(),
			})
			return
		}

		newC, err := couponAPI.CreateCoupon(ctx.Context(), tit, exp, dis)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(201).JSON(newC)
	}
}
