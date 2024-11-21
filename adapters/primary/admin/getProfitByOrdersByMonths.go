package admins

import (
	orders "github.com/ecommerce/core/services/order"
	"github.com/gofiber/fiber/v2"
)

func GetProfitByOrdersByMonths(orderAPI orders.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ords, errO := orderAPI.GetProfitByOrdersByMonths(ctx.Context())
		if errO != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errO.Error(),
			})
		}

		return ctx.Status(200).JSON(ords)
	}
}
