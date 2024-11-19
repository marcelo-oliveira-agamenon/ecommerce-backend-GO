package favorites

import (
	"errors"

	favorites "github.com/ecommerce/core/services/favorite"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingFavoriteIdField = errors.New("missing favorite id parameter")
	ErrorDeletingFavorite       = errors.New("unknown deleting field")
)

func DeleteFavorite(favoriteAPI favorites.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		favId := ctx.Params("id")
		if favId == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingFavoriteIdField.Error(),
			})
		}

		isDel, errF := favoriteAPI.DeleteFavorite(ctx.Context(), favId)
		if errF != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errF.Error(),
			})
		}
		if !isDel {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorDeletingFavorite.Error(),
			})
		}

		ctx.Status(200)
		return nil
	}
}
