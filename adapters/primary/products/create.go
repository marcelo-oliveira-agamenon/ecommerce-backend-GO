package products

import (
	"errors"
	"strconv"

	"github.com/ecommerce/core/domain/product"
	categories "github.com/ecommerce/core/services/category"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/core/util"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	AuthHeader           = "Authorization"
	ErrorConversion      = errors.New("conversion function failed")
	ErrorCategoryUnknown = errors.New("unknown category")
)

func CreateProduct(productAPI products.API, categoriesAPI categories.API, token ports.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		tok, errT := util.GetToken(ctx, AuthHeader)
		if errT != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": errT.Error(),
			})
			return
		}

		if err := ctx.BodyParser(&product.Product{}); err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		_, errC := token.ClaimTokenData(*tok)
		if errC != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": errC.Error(),
			})
			return
		}

		catId := ctx.FormValue("categoryid")
		_, errCat := categoriesAPI.GetCategoryById(ctx.Context(), catId)
		if errCat != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorCategoryUnknown.Error(),
			})
			return
		}

		value, e1 := strconv.ParseFloat(ctx.FormValue("value"), 64)
		stockQtd, e2 := strconv.Atoi(ctx.FormValue("stockqtd"))
		hasPromo, e3 := strconv.ParseBool(ctx.FormValue("hasPromotion"))
		hasShip, e4 := strconv.ParseBool(ctx.FormValue("hasShipping"))
		discount, e5 := strconv.ParseFloat(ctx.FormValue("discount"), 64)
		price, e6 := strconv.ParseFloat(ctx.FormValue("shippingPrice"), 64)
		if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil || e6 != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorConversion.Error(),
			})
			return
		}

		product, errP := productAPI.CreateProduct(ctx.Context(), product.Product{
			Name:            ctx.FormValue("name"),
			Description:     ctx.FormValue("description"),
			TypeUnit:        ctx.FormValue("type"),
			TecnicalDetails: ctx.FormValue("tecnicalDetails"),
			Categoryid:      catId,
			Value:           value,
			StockQtd:        stockQtd,
			HasPromotion:    hasPromo,
			HasShipping:     hasShip,
			Discount:        discount,
			ShippingPrice:   price,
			Rate:            0,
		})
		if errP != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
			return
		}

		ctx.Status(201).JSON(product)
	}
}
