package payments

import (
	"github.com/ecommerce/core/services/payments"
	"github.com/gofiber/fiber/v2"
)

func DeletePayment(paymentAPI payments.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		pde, errD := paymentAPI.DeletePayment(ctx.Context(), id)
		if errD != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errD.Error(),
			})
		}

		return ctx.Status(200).JSON(pde)
	}
}
