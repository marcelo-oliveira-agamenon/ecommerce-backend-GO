package reports

import (
	"errors"

	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingCreatedAtStart = errors.New("missing created at start in query parameters")
	ErrorMissingCreatedAtEnd   = errors.New("missing created at end in query parameters")
	ErrorMissingGender         = errors.New("missing gender in query parameters")
)

func ExportUsers(userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		cAtSt := ctx.Query("user_created_at_start")
		cAtEn := ctx.Query("user_created_at_end")
		gen := ctx.Query("gender")
		if cAtSt == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingCreatedAtStart.Error(),
			})
		}
		if cAtEn == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingCreatedAtEnd.Error(),
			})
		}
		if gen == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingGender.Error(),
			})
		}

		expU, errU := userAPI.ExportUsers(ctx.Context(), cAtSt, cAtEn, gen)
		if errU != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errU.Error(),
			})
		}

		return ctx.Status(200).JSON(expU)
	}
}
