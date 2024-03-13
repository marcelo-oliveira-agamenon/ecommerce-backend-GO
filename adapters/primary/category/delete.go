package categories

import (
	"errors"

	categories "github.com/ecommerce/core/services/category"
	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingIdParams     = errors.New("missing id params")
	ErrorCategoryDoenstExist = errors.New("this category doenst exist")
)

func DeleteCategory(categoryAPI categories.API, productsAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		catId := ctx.Params("id")
		if catId == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingIdParams.Error(),
			})
			return
		}

		cat, errG := categoryAPI.GetCategoryById(ctx.Context(), catId)
		if errG != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorCategoryDoenstExist.Error(),
			})
			return
		}

		_, errA := categoryAPI.DeleteCategory(ctx.Context(), *cat)
		if errA != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errA.Error(),
			})
			return
		}

		ctx.Status(200)
	}
}
