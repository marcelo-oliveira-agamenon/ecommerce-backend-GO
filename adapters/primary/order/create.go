package orders

import (
	"errors"

	"github.com/ecommerce/core/domain/orderDetails"
	orders "github.com/ecommerce/core/services/order"
	ordersdetails "github.com/ecommerce/core/services/ordersDetails"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorInvalidOrderData = errors.New("invalid order data")
)

func CreateOrder(orderAPI orders.API, userAPI users.API, orderDetailsAPI ordersdetails.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dec := ctx.Locals("user").(*ports.Claims)

		_, errU := userAPI.GetUserById(ctx.Context(), dec.UserId)
		if errU != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": errU.Error(),
			})
		}

		var orDe []orderDetails.OrderProductData
		errPr := ctx.BodyParser(&orDe)
		if errPr != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorInvalidOrderData.Error(),
			})
		}

		var ordSli []orderDetails.OrderDetails
		for _, val := range orDe {
			orFor, errF := orderDetailsAPI.CheckOrderDetails(ctx.Context(), "placeholder", val)
			if errF != nil {
				return ctx.Status(422).JSON(&fiber.Map{
					"error": errF.Error(),
				})
			}

			ordSli = append(ordSli, *orFor)
		}

		toV, qtd, errDe := orderDetailsAPI.GetTotalOrderValueAndQuatity(ctx.Context(), orDe)
		if errDe != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errDe.Error(),
			})
		}

		newO, err := orderAPI.AddOrder(ctx.Context(), dec.UserId, *qtd, *toV, ordSli, orDe[0].CouponUsed)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		//TODO: send email about order

		return ctx.Status(201).JSON(newO)
	}
}
