package categories

import (
	"errors"

	categories "github.com/ecommerce/core/services/category"
	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingIdParams     = errors.New("missing id params")
	ErrorCategoryDoenstExist = errors.New("this category doenst exist")
)

func DeleteCategory(categoryAPI categories.API, productsAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		catId := ctx.Params("id")
		if catId == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingIdParams.Error(),
			})
		}

		cat, errG := categoryAPI.GetCategoryById(ctx.Context(), catId)
		if errG != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorCategoryDoenstExist.Error(),
			})
		}

		_, errA := categoryAPI.DeleteCategory(ctx.Context(), *cat)
		if errA != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errA.Error(),
			})
		}

		ctx.Status(200)
		return nil
	}
}
