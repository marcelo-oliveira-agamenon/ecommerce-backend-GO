package orders

import (
	"errors"
	"strconv"

	orders "github.com/ecommerce/core/services/order"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	AuthHeader              = "Authorization"
	ErrorMissingOffsetLimit = errors.New("missing limit or offset query parameter")
)

func GetByUserId(orderAPI orders.API, userAPI users.API) fiber.Handler {
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

		ords, err := orderAPI.GetByUserId(ctx.Context(), dec.UserId, limit, offset)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(200).JSON(ords)
	}
}
