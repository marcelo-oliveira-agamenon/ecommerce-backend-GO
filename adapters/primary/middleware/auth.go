package middleware

import (
	"github.com/ecommerce/core/util"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	AuthHeader = "Authorization"
)

func VerifyToken(j ports.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token, errTo := util.GetToken(ctx, AuthHeader)
		if errTo != nil {
			return ctx.Status(401).JSON(&fiber.Map{
				"error": errTo.Error(),
			})
		}

		err := j.VerifyToken(*token)
		if err != nil {
			return ctx.Status(401).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		dec, errC := j.ClaimTokenData(*token)
		if errC != nil {
			return ctx.Status(401).JSON(&fiber.Map{
				"error": errC.Error(),
			})
		}

		ctx.Locals("user", dec)
		return ctx.Next()
	}
}
