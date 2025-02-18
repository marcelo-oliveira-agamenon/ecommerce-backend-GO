package products

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ecommerce/core/domain/product"
	categories "github.com/ecommerce/core/services/category"
	logs "github.com/ecommerce/core/services/log"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

var (
	AuthHeader           = "Authorization"
	ErrorConversion      = errors.New("fields missing/incorrect")
	ErrorCategoryUnknown = errors.New("unknown category")
)

func CreateProduct(productAPI products.API, categoriesAPI categories.API, logAPI logs.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.BodyParser(&product.Product{}); err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		dec := ctx.Locals("user").(*ports.Claims)

		catId := ctx.FormValue("categoryid")
		_, errCat := categoriesAPI.GetCategoryById(ctx.Context(), catId)
		if errCat != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorCategoryUnknown.Error(),
			})
		}

		value, e1 := strconv.ParseFloat(ctx.FormValue("value"), 64)
		stockQtd, e2 := strconv.Atoi(ctx.FormValue("stockqtd"))
		hasPromo, e3 := strconv.ParseBool(ctx.FormValue("hasPromotion"))
		hasShip, e4 := strconv.ParseBool(ctx.FormValue("hasShipping"))
		discount, e5 := strconv.ParseFloat(ctx.FormValue("discount"), 64)
		price, e6 := strconv.ParseFloat(ctx.FormValue("shippingPrice"), 64)
		if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil || e6 != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorConversion.Error(),
			})
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
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
		}

		msg := "Insert product with ID: " + product.ID
		errL := logAPI.AddLog(ctx.Context(), "product", msg, uuid.FromStringOrNil(dec.UserId))
		if errL != nil {
			fmt.Println(errL)
		}

		return ctx.Status(201).JSON(product)
	}
}
