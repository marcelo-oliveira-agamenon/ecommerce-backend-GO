package products

import (
	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber/v2"
)

func DeleteProductById(productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		prodId := ctx.Params("id")

		prod, errG := productAPI.GetProductById(ctx.Context(), prodId)
		if errG != nil {
			return ctx.Status(404).JSON(&fiber.Map{
				"error": errG.Error(),
			})
		}

		errD := productAPI.DeleteProductById(ctx.Context(), *prod)
		if errD != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errD.Error(),
			})
		}

		ctx.Status(200)
		return nil
	}
}
