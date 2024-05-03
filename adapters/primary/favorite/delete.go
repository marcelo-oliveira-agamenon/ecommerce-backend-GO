package favorites

import (
	"errors"

	favorites "github.com/ecommerce/core/services/favorite"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingFavoriteIdField = errors.New("missing favorite id parameter")
	ErrorDeletingFavorite       = errors.New("unknown deleting field")
)

func DeleteFavorite(favoriteAPI favorites.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		favId := ctx.Params("id")
		if favId == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingFavoriteIdField.Error(),
			})
			return
		}

		isDel, errF := favoriteAPI.DeleteFavorite(ctx.Context(), favId)
		if errF != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errF.Error(),
			})
			return
		}
		if !isDel {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorDeletingFavorite.Error(),
			})
			return
		}

		ctx.Status(200)
	}
}
