package orders

import (
	"errors"

	orders "github.com/ecommerce/core/services/order"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingRateParams = errors.New("missing rate parameter")
)

func EditRate(orderAPI orders.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		rate := ctx.Query("rate")
		if rate == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingRateParams.Error(),
			})
		}

		or, err := orderAPI.UpdateRate(ctx.Context(), id, rate)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(200).JSON(or)
	}
}
