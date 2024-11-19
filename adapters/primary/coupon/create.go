package coupons

import (
	"errors"

	coupons "github.com/ecommerce/core/services/coupon"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingExpireDate  = errors.New("missing expire date")
	ErrorMissingDiscount    = errors.New("missing discount")
	ErrorMissingTitle       = errors.New("title")
	ErrorCouponAlreadyExist = errors.New("coupon already exist with this title")
)

func CreateCoupon(couponAPI coupons.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		exp := ctx.FormValue("expire_date")
		if exp == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingExpireDate.Error(),
			})
		}

		dis := ctx.FormValue("discount")
		if dis == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingDiscount.Error(),
			})
		}

		tit := ctx.FormValue("title")
		if tit == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingTitle.Error(),
			})
		}

		hasCoupon, errH := couponAPI.CheckIfThereIsCouponsByHash(ctx.Context(), tit)
		if errH != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errH.Error(),
			})
		}
		if hasCoupon {
			return ctx.Status(409).JSON(&fiber.Map{
				"error": ErrorCouponAlreadyExist.Error(),
			})
		}

		newC, err := couponAPI.CreateCoupon(ctx.Context(), tit, exp, dis)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(201).JSON(newC)
	}
}
