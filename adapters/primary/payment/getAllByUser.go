package payments

import (
	"errors"
	"strconv"

	"github.com/ecommerce/core/services/payments"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	AuthHeader              = "Authorization"
	ErrorMissingOffsetLimit = errors.New("missing limit or offset query parameter")
)

func GetAllByUser(paymentAPI payments.API, userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		dec := ctx.Locals("user").(*ports.Claims)

		_, errU := userAPI.GetUserById(ctx.Context(), dec.UserId)
		if errU != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": errU.Error(),
			})
			return
		}

		limit, err1 := strconv.Atoi(ctx.Query("limit"))
		offset, err2 := strconv.Atoi(ctx.Query("offset"))
		if err1 != nil || err2 != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingOffsetLimit.Error(),
			})
			return
		}

		list, errG := paymentAPI.GetAllPaymentsByUser(ctx.Context(), dec.UserId, limit, offset)
		if errG != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
			return
		}

		ctx.Status(200).JSON(list)
	}
}
