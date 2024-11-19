package products

import (
	"errors"
	"strconv"

	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingOffsetLimit = errors.New("missing limit or offset query parameter")
)

func GetAllProducts(productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		limit, err1 := strconv.Atoi(ctx.Query("limit"))
		offset, err2 := strconv.Atoi(ctx.Query("offset"))
		if err1 != nil || err2 != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingOffsetLimit.Error(),
			})
		}

		getByCategory := ctx.Query("category")
		getByPromotion := ctx.Query("promotion")
		getRecentOnes := ctx.Query("recent")
		getByName := ctx.Query("name")

		list, count, errG := productAPI.GetAllProducts(ctx.Context(), limit, offset, getByCategory, getByPromotion, getRecentOnes, getByName)
		if errG != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
		}

		return ctx.Status(200).JSON(&fiber.Map{
			"products": list,
			"count":    count,
		})
	}
}
