package admins

import (
	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber"
)

func GetProductQuantityByCategories(productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		prod, tot, errP := productAPI.GetProductQuantityByCategories(ctx.Context())
		if errP != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"errPr": errP.Error(),
			})
			return
		}

		ctx.Status(200).JSON(&fiber.Map{
			"data":  prod,
			"total": tot,
		})
	}
}
