package users

import (
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/core/util"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	AuthHeader = "Authorization"
)

func RefreshToken(userAPI users.API, token ports.TokenService, redis ports.RedisService) fiber.Handler {
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

		errV := redis.ValidateSession(ctx.Context(), dec.UserId)
		if errV != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errV.Error(),
			})
			return
		}

		token, exTi, errToken := token.CreateToken(dec.UserId)
		if errToken != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errToken.Error(),
			})
			return
		}

		errR := redis.StoreUserSession(ctx.Context(), dec.UserId, exTi)
		if errR != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errR.Error(),
			})
			return
		}

		ctx.Status(200).JSON(&fiber.Map{
			"token": token,
		})
	}
}
