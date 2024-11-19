package categories

import (
	"errors"
	"strconv"

	categories "github.com/ecommerce/core/services/category"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingOffsetLimit = errors.New("missing limit or offset query parameter")
)

func GetAllCategories(categoryAPI categories.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		limit, err1 := strconv.Atoi(ctx.Query("limit"))
		offset, err2 := strconv.Atoi(ctx.Query("offset"))
		if err1 != nil || err2 != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingOffsetLimit.Error(),
			})
		}

		list, err := categoryAPI.GetAllCategories(ctx.Context(), limit, offset)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(200).JSON(list)
	}
}
