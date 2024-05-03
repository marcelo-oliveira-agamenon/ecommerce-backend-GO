package payments

import (
	"errors"

	orders "github.com/ecommerce/core/services/order"
	"github.com/ecommerce/core/services/payments"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/core/util"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingOrderId       = errors.New("missing order id in body parameters")
	ErrorMissingTotalValue    = errors.New("missing total value in body parameters")
	ErrorMissingPaidValue     = errors.New("missing paid value in body parameters")
	ErrorMissingPaymentMethod = errors.New("missing payment method in body parameters")
)

func CreatePayment(paymentAPI payments.API, userAPI users.API, orderAPI orders.API, token ports.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		tok, errT := util.GetToken(ctx, AuthHeader)
		if errT != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": errT.Error(),
			})
			return
		}

		dec, errC := token.ClaimTokenData(*tok)
		if errC != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": errC.Error(),
			})
			return
		}

		_, errU := userAPI.GetUserById(ctx.Context(), dec.UserId)
		if errU != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": errU.Error(),
			})
			return
		}

		orId := ctx.FormValue("order_id")
		toVa := ctx.FormValue("total_value")
		paV := ctx.FormValue("paid_value")
		paM := ctx.FormValue("payment_method")
		if orId == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingOrderId.Error(),
			})
			return
		}
		if toVa == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingTotalValue.Error(),
			})
			return
		}
		if paV == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingPaidValue.Error(),
			})
			return
		}
		if paM == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingPaymentMethod.Error(),
			})
			return
		}

		_, errO := orderAPI.GetById(ctx.Context(), orId)
		if errO != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": errO.Error(),
			})
			return
		}

		nPay, errG := paymentAPI.CreatePayment(ctx.Context(), dec.UserId, orId, toVa, paV, paM)
		if errG != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
			return
		}

		//TODO: add to log table

		ctx.Status(201).JSON(nPay)
	}
}
