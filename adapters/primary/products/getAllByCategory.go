package products

import (
	"strconv"

	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber/v2"
)

func GetAllProductsByCategory(productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorParameter.Error(),
			})
		}

		limit, err1 := strconv.Atoi(ctx.Query("limit"))
		offset, err2 := strconv.Atoi(ctx.Query("offset"))
		if err1 != nil || err2 != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingOffsetLimit.Error(),
			})
		}

		list, errG := productAPI.GetProductsByCategory(ctx.Context(), id, limit, offset)
		if errG != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
		}

		return ctx.Status(200).JSON(list)
	}
}
