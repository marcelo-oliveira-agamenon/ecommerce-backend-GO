package users

import (
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

func LoginFacebook(userAPI users.API, token ports.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) {

	}
}
