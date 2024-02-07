package products

import (
	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber"
)

func DeleteProductById(productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		prodId := ctx.Params("id")

		prod, errG := productAPI.GetProductById(ctx.Context(), prodId)
		if errG != nil {
			ctx.Status(404).JSON(&fiber.Map{
				"error": errG.Error(),
			})
			return
		}

		errD := productAPI.DeleteProductById(ctx.Context(), *prod)
		if errD != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errD.Error(),
			})
			return
		}

		ctx.Status(200)
	}
}
