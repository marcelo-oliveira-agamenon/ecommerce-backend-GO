package users

import (
	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber"
)

func ResetPassword(userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		rePas := new(users.ResetPassword)
		ctx.BodyParser(rePas)

		_, err := userAPI.ResetPassword(ctx.Context(), *rePas)
		if err != nil {

			return
		}
	}
}
