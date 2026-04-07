package users

import (
	logs "github.com/ecommerce/core/services/log"
	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func ToggleRoles(userAPI users.API, logAPI logs.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Params("id")

		user, err := userAPI.ToggleRoles(ctx.Context(), userId)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		msg := "Role of user with ID: " + userId + " was changed"
		errA := logAPI.AddLog(ctx.Context(), "user", msg, uuid.FromStringOrNil(userId))
		if errA != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errA.Error(),
			})
		}

		return ctx.Status(200).JSON(user)
	}
}
