package users

import (
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

func Login(userAPI users.API, token ports.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		lgReq := new(users.LoginRequest)
		ctx.BodyParser(lgReq)
		lgReq.IsAdmin = ctx.Query("admin")

		user, errU := userAPI.Login(ctx.Context(), *lgReq)
		if errU != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errU.Error(),
			})
			return
		}

		token, _, errToken := token.CreateToken(user.ID.String())
		if errToken != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errToken.Error(),
			})
			return
		}

		ctx.Status(200).JSON(&fiber.Map{
			"user":  user,
			"token": token,
		})
	}
}
