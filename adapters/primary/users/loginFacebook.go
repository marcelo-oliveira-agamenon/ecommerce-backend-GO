package users

import (
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

func LoginFacebook(userAPI users.API, token ports.TokenService, redis ports.RedisService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		lgReq := new(users.LoginFacebook)
		ctx.BodyParser(lgReq)

		user, errU := userAPI.LoginFacebook(ctx.Context(), *lgReq)
		if errU != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errU.Error(),
			})
			return
		}

		token, exTi, errToken := token.CreateToken(user.ID.String())
		if errToken != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errToken.Error(),
			})
			return
		}

		errR := redis.StoreUserSession(ctx.Context(), user.ID.String(), exTi)
		if errR != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errR.Error(),
			})
			return
		}

		ctx.Status(200).JSON(&fiber.Map{
			"user":  user,
			"token": token,
		})
	}
}
