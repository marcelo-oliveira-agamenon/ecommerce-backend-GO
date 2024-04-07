package payments

import (
	"github.com/ecommerce/core/services/payments"
	"github.com/gofiber/fiber"
)

func DeletePayment(paymentAPI payments.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		pde, errD := paymentAPI.DeletePayment(ctx.Context(), id)
		if errD != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errD.Error(),
			})
			return
		}

		ctx.Status(200).JSON(pde)
	}
}
