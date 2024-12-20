package products

import (
	"errors"

	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorParameter = errors.New("missing id parameter")
)

func GetProductById(productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorParameter.Error(),
			})
		}

		prod, errG := productAPI.GetProductById(ctx.Context(), id)
		if errG != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
		}

		return ctx.Status(200).JSON(prod)
	}
}
