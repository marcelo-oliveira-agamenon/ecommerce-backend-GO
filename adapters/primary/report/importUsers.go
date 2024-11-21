package reports

import (
	"errors"

	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorCsvFileParam = errors.New("reading csv file")
	ErrorCsvOpenFile  = errors.New("open csv file")
	ErrorAddUsers     = errors.New("importing users from file")
)

func ImportUsers(userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		file, errF := ctx.FormFile("csv_file")
		if errF != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorCsvFileParam.Error(),
			})
		}

		fiOp, errO := file.Open()
		if errO != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorCsvOpenFile.Error(),
			})
		}

		expU, errU := userAPI.ImportUsers(ctx.Context(), fiOp)
		if errU != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errU.Error(),
			})
		}

		if expU {
			ctx.Status(200)
			return nil
		}

		return ctx.Status(500).JSON(&fiber.Map{
			"error": ErrorAddUsers.Error(),
		})
	}
}
