package payments

import (
	"errors"

	orders "github.com/ecommerce/core/services/order"
	"github.com/ecommerce/core/services/payments"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingOrderId       = errors.New("missing order id in body parameters")
	ErrorMissingTotalValue    = errors.New("missing total value in body parameters")
	ErrorMissingPaidValue     = errors.New("missing paid value in body parameters")
	ErrorMissingPaymentMethod = errors.New("missing payment method in body parameters")
)

func CreatePayment(paymentAPI payments.API, userAPI users.API, orderAPI orders.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dec := ctx.Locals("user").(*ports.Claims)

		_, errU := userAPI.GetUserById(ctx.Context(), dec.UserId)
		if errU != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": errU.Error(),
			})
		}

		orId := ctx.FormValue("order_id")
		toVa := ctx.FormValue("total_value")
		paV := ctx.FormValue("paid_value")
		paM := ctx.FormValue("payment_method")
		if orId == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingOrderId.Error(),
			})
		}
		if toVa == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingTotalValue.Error(),
			})
		}
		if paV == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingPaidValue.Error(),
			})
		}
		if paM == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingPaymentMethod.Error(),
			})
		}

		_, errO := orderAPI.GetById(ctx.Context(), orId)
		if errO != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": errO.Error(),
			})
		}

		nPay, errG := paymentAPI.CreatePayment(ctx.Context(), dec.UserId, orId, toVa, paV, paM)
		if errG != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
		}

		//TODO: add to log table

		return ctx.Status(201).JSON(nPay)
	}
}
