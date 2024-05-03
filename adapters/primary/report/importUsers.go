package reports

import (
	"errors"

	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber"
)

var (
	ErrorCsvFileParam = errors.New("reading csv file")
	ErrorCsvOpenFile  = errors.New("open csv file")
	ErrorAddUsers     = errors.New("importing users from file")
)

func ImportUsers(userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		file, errF := ctx.FormFile("csv_file")
		if errF != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorCsvFileParam.Error(),
			})
			return
		}

		fiOp, errO := file.Open()
		if errO != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorCsvOpenFile.Error(),
			})
			return
		}

		expU, errU := userAPI.ImportUsers(ctx.Context(), fiOp)
		if errU != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errU.Error(),
			})
			return
		}

		if expU {
			ctx.Status(200)
			return
		}

		ctx.Status(500).JSON(&fiber.Map{
			"error": ErrorAddUsers.Error(),
		})
		return
	}
}
