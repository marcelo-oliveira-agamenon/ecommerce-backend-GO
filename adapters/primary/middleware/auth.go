package middleware

import (
	"errors"
	"strings"

	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingToken = errors.New("missing token")
)

func VerifyToken(j ports.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		rawToken := strings.Replace(string(ctx.Fasthttp.Request.Header.Peek("Authorization")), "Bearer ", "", 1)
		if rawToken == "" || rawToken == "null" {
			ctx.Status(401).JSON(&fiber.Map{
				"error": ErrorMissingToken.Error(),
			})
			return
		}

		inv := j.VerifyToken(rawToken)
		if inv != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": inv.Error(),
			})
			return
		}

		ctx.Next()
	}
}
