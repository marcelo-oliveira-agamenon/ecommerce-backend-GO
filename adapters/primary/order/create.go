package orders

import (
	"errors"
	"strconv"

	orders "github.com/ecommerce/core/services/order"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingProductIds = errors.New("missing product id")
	ErrorMissingQuantity   = errors.New("missing quantity")
	ErrorMissingTotalValue = errors.New("missing total value")
)

func CreateOrder(orderAPI orders.API, userAPI users.API, productAPI products.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		dec := ctx.Locals("user").(*ports.Claims)

		_, errU := userAPI.GetUserById(ctx.Context(), dec.UserId)
		if errU != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": errU.Error(),
			})
			return
		}

		prodId := ctx.FormValue("productID")
		if len(prodId) == 0 {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingProductIds.Error(),
			})
			return
		}

		qtd, err1 := strconv.Atoi(ctx.FormValue("qtd"))
		toV, err2 := strconv.ParseFloat(ctx.FormValue("totalValue"), 64)
		if err1 != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingQuantity.Error(),
			})
			return
		}
		if err2 != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingTotalValue.Error(),
			})
			return
		}

		_, errP := productAPI.CheckProductListById(ctx.Context(), prodId)
		if errP != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
			return
		}

		newO, err := orderAPI.AddOrder(ctx.Context(), dec.UserId, prodId, qtd, toV)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		//TODO: send email about order

		ctx.Status(201).JSON(newO)
	}
}
