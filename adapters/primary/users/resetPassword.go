package users

import (
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	PasswordChangedMessage = "Password changed"
)

func ResetPassword(userAPI users.API, redis ports.RedisService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		rePas := new(users.ResetPassword)
		if errC := ctx.BodyParser(rePas); errC != nil {
			return ctx.Status(400).JSON(&fiber.Map{
				"error": errC.Error(),
			})
		}

		errR := redis.ValidateResetPasswordInfo(ctx.Context(), rePas.Hash)
		if errR != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errR.Error(),
			})
		}

		_, err := userAPI.ResetPassword(ctx.Context(), *rePas)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(200).JSON(PasswordChangedMessage)
	}
}
