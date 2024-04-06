package orders

import (
	"errors"

	orders "github.com/ecommerce/core/services/order"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingPaidParams = errors.New("missing paid parameter")
)

func EditPayment(orderAPI orders.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		paid := ctx.Query("paid")
		if paid == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingPaidParams.Error(),
			})
			return
		}

		or, err := orderAPI.UpdatePayment(ctx.Context(), id, paid)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(200).JSON(or)
	}
}
