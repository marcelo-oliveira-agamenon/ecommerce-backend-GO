package admins

import (
	orders "github.com/ecommerce/core/services/order"
	"github.com/gofiber/fiber/v2"
)

func GetOrdersByPeriod(orderAPI orders.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ordP, errO := orderAPI.GetOrdersByPeriod(ctx.Context())
		if errO != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errO.Error(),
			})

		}

		return ctx.Status(200).JSON(ordP)
	}
}
