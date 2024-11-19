package orders

import (
	"errors"
	"strconv"

	orders "github.com/ecommerce/core/services/order"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingProductIds = errors.New("missing product id")
	ErrorMissingQuantity   = errors.New("missing quantity")
	ErrorMissingTotalValue = errors.New("missing total value")
)

func CreateOrder(orderAPI orders.API, userAPI users.API, productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dec := ctx.Locals("user").(*ports.Claims)

		_, errU := userAPI.GetUserById(ctx.Context(), dec.UserId)
		if errU != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": errU.Error(),
			})
		}

		prodId := ctx.FormValue("productID")
		if len(prodId) == 0 {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingProductIds.Error(),
			})
		}

		qtd, err1 := strconv.Atoi(ctx.FormValue("qtd"))
		toV, err2 := strconv.ParseFloat(ctx.FormValue("totalValue"), 64)
		if err1 != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingQuantity.Error(),
			})
		}
		if err2 != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingTotalValue.Error(),
			})
		}

		_, errP := productAPI.CheckProductListById(ctx.Context(), prodId)
		if errP != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
		}

		newO, err := orderAPI.AddOrder(ctx.Context(), dec.UserId, prodId, qtd, toV)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		//TODO: send email about order

		return ctx.Status(201).JSON(newO)
	}
}
