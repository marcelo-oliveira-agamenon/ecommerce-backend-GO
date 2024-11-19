package admins

import (
	orders "github.com/ecommerce/core/services/order"
	products "github.com/ecommerce/core/services/products"
	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber/v2"
)

func GetCountForAdmin(userAPI users.API, orderAPI orders.API, productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		coUs, errU := userAPI.GetUserCount(ctx.Context())
		if errU != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errU.Error(),
			})
		}

		coPr, errP := productAPI.GetProductCount(ctx.Context())
		if errP != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
		}

		coOA, coOP, errP := orderAPI.GetOrderCount(ctx.Context())
		if errP != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
		}

		return ctx.Status(200).JSON(&fiber.Map{
			"countAllOrders":     coOA,
			"countAllPaidOrders": coOP,
			"countAllProducts":   coPr,
			"countAllUsers":      coUs,
		})
	}
}
