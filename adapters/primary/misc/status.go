package miscs

import (
	"github.com/ecommerce/core/services/misc"
	"github.com/gofiber/fiber/v2"
)

func DatabaseStatus(miscAPI misc.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		status := miscAPI.GetDatabaseStatus()

		return ctx.Status(200).JSON(status)
	}
}
