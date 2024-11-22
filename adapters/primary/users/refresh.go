package users

import (
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	AuthHeader = "Authorization"
)

func RefreshToken(token ports.TokenService, redis ports.RedisService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dec := ctx.Locals("user").(*ports.Claims)

		errV := redis.ValidateSession(ctx.Context(), dec.UserId)
		if errV != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errV.Error(),
			})
		}

		token, exTi, errToken := token.CreateToken(dec.UserId)
		if errToken != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errToken.Error(),
			})
		}

		errR := redis.StoreUserSession(ctx.Context(), dec.UserId, exTi, ctx.IP(), ctx.Get("User-Agent"))
		if errR != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errR.Error(),
			})
		}

		return ctx.Status(200).JSON(&fiber.Map{
			"token": token,
		})
	}
}
