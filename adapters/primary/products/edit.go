package products

import (
	"strconv"

	"github.com/ecommerce/core/domain/product"
	"github.com/ecommerce/core/services/products"
	"github.com/gofiber/fiber/v2"
)

func EditProduct(productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.BodyParser(&product.Product{}); err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		prodId := ctx.Params("id")
		_, errI := productAPI.GetProductById(ctx.Context(), prodId)
		if errI != nil {
			return ctx.Status(404).JSON(&fiber.Map{
				"error": errI.Error(),
			})
		}

		value, e1 := strconv.ParseFloat(ctx.FormValue("value"), 64)
		stockQtd, e2 := strconv.Atoi(ctx.FormValue("stockqtd"))
		hasPromo, e3 := strconv.ParseBool(ctx.FormValue("hasPromotion"))
		hasShip, e4 := strconv.ParseBool(ctx.FormValue("hasShipping"))
		discount, e5 := strconv.ParseFloat(ctx.FormValue("discount"), 64)
		price, e6 := strconv.ParseFloat(ctx.FormValue("shippingPrice"), 64)
		if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil || e6 != nil {
			return ctx.Status(404).JSON(&fiber.Map{
				"error": ErrorConversion.Error(),
			})
		}

		product, errP := productAPI.EditProduct(ctx.Context(), product.Product{
			Name:            ctx.FormValue("name"),
			Description:     ctx.FormValue("description"),
			TypeUnit:        ctx.FormValue("type"),
			TecnicalDetails: ctx.FormValue("tecnicalDetails"),
			Categoryid:      ctx.FormValue("categoryid"),
			Value:           value,
			StockQtd:        stockQtd,
			HasPromotion:    hasPromo,
			HasShipping:     hasShip,
			Discount:        discount,
			ShippingPrice:   price,
			Rate:            0,
		})
		if errP != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
		}

		return ctx.Status(200).JSON(product)
	}
}
