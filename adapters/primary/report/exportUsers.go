package reports

import (
	"errors"

	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingCreatedAtStart = errors.New("missing created at start in query parameters")
	ErrorMissingCreatedAtEnd   = errors.New("missing created at end in query parameters")
	ErrorMissingGender         = errors.New("missing gender in query parameters")
)

func ExportUsers(userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		cAtSt := ctx.Query("user_created_at_start")
		cAtEn := ctx.Query("user_created_at_end")
		gen := ctx.Query("gender")
		if cAtSt == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingCreatedAtStart.Error(),
			})
			return
		}
		if cAtEn == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingCreatedAtEnd.Error(),
			})
			return
		}
		if gen == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingGender.Error(),
			})
			return
		}

		expU, errU := userAPI.ExportUsers(ctx.Context(), cAtSt, cAtEn, gen)
		if errU != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errU.Error(),
			})
			return
		}

		ctx.Status(200).JSON(expU)
	}
}
