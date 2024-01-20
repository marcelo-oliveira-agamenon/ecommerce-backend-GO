package middleware

import (
	"github.com/ecommerce/core/util"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	AuthHeader = "Authorization"
)

func VerifyToken(j ports.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		token, errTo := util.GetToken(ctx, AuthHeader)
		if errTo != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": errTo.Error(),
			})
			return
		}

		err := j.VerifyToken(*token)
		if err != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.Next()
	}
}
