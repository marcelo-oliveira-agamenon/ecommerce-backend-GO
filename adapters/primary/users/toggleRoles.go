package users

import (
	logs "github.com/ecommerce/core/services/log"
	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
)

func ToggleRoles(userAPI users.API, logAPI logs.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		userId := ctx.Params("id")

		user, err := userAPI.ToggleRoles(ctx.Context(), userId)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		msg := "Role of user with ID: " + userId + " was changed"
		logAPI.AddLog(ctx.Context(), "user", msg, uuid.FromStringOrNil(userId))

		ctx.Status(200).JSON(user)
	}
}
