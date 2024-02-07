package categories

import (
	categories "github.com/ecommerce/core/services/category"
	"github.com/gofiber/fiber"
)

func GetAllCategories(categoryAPI categories.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		list, err := categoryAPI.GetAllCategories(ctx.Context())
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(200).JSON(list)
	}
}
