package middleware

import (
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

func VerifyAdminPermission(us users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dec := ctx.Locals("user").(*ports.Claims)

		user, errU := us.GetUserById(ctx.Context(), dec.UserId)
		if errU != nil {
			return ctx.Status(500).JSON(errU)
		}

		isAdmin := false
		for _, v := range user.Roles {
			if v == "admin" {
				isAdmin = true
			}
		}

		if !isAdmin {
			return ctx.Status(401).JSON("You dont have permission to this action")
		}

		return ctx.Next()
	}
}
