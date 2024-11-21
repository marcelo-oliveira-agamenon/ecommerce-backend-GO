package users

import (
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

func Logout(redis ports.RedisService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*ports.Claims)

		errR := redis.ClearUserSession(ctx.Context(), user.UserId)
		if errR != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errR.Error(),
			})
		}

		ctx.Status(200)
		return nil
	}
}
