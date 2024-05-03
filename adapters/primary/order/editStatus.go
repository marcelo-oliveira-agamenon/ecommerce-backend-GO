package orders

import (
	"errors"

	orders "github.com/ecommerce/core/services/order"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingStatusParams = errors.New("missing status parameter")
)

func EditStatus(orderAPI orders.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		paid := ctx.Query("status")
		if paid == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingStatusParams.Error(),
			})
			return
		}

		or, err := orderAPI.UpdateStatus(ctx.Context(), id, paid)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(200).JSON(or)
	}
}
