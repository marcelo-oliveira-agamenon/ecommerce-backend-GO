package users

import (
	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber"
)

func ToggleRoles(userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		userId := ctx.Params("id")

		user, err := userAPI.ToggleRoles(ctx.Context(), userId)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(200).JSON(user)
	}
}
