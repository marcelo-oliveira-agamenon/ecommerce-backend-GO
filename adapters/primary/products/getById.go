package products

import (
	"errors"

	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber"
)

var (
	ErrorParameter = errors.New("missing id parameter")
)

func GetProductById(productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		if id == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorParameter,
			})
			return
		}

		prod, errG := productAPI.GetProductById(ctx.Context(), id)
		if errG != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
			return
		}

		ctx.Status(200).JSON(prod)
	}
}
