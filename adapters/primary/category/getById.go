package categories

import (
	"errors"

	categories "github.com/ecommerce/core/services/category"
	"github.com/gofiber/fiber"
)

var (
	ErrorParameter = errors.New("missing id parameter")
)

func GetOneCategory(categoryAPI categories.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		if id == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorParameter.Error(),
			})
			return
		}

		cat, err := categoryAPI.GetCategoryById(ctx.Context(), id)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(200).JSON(cat)
	}
}
