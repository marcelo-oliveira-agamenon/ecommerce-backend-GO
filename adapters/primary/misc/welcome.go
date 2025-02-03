package miscs

import (
	"github.com/gofiber/fiber/v2"
)

func WelcomeAPIReturn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("Welcome to Cash and Grab API")
	}
}
