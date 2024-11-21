package orders

import (
	"errors"

	orders "github.com/ecommerce/core/services/order"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingStatusParams = errors.New("missing status parameter")
)

func EditStatus(orderAPI orders.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		paid := ctx.Query("status")
		if paid == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingStatusParams.Error(),
			})
		}

		or, err := orderAPI.UpdateStatus(ctx.Context(), id, paid)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(200).JSON(or)
	}
}
