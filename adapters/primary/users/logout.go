package users

import (
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

func Logout(redis ports.RedisService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		user := ctx.Locals("user").(*ports.Claims)

		errR := redis.ClearUserSession(ctx.Context(), user.UserId)
		if errR != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errR.Error(),
			})
			return
		}

		ctx.Status(200)
	}
}
