package users

import (
	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber/v2"
)

var (
	PasswordChangedMessage = "Password changed"
)

func ResetPassword(userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		rePas := new(users.ResetPassword)
		ctx.BodyParser(rePas)

		_, err := userAPI.ResetPassword(ctx.Context(), *rePas)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(200).JSON(PasswordChangedMessage)
	}
}
