package favorites

import (
	"errors"
	"strconv"

	favorites "github.com/ecommerce/core/services/favorite"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingOffsetLimit = errors.New("missing limit or offset query parameter")
)

func GetFavoriteByUserId(favoriteAPI favorites.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		dec := ctx.Locals("user").(*ports.Claims)

		limit, err1 := strconv.Atoi(ctx.Query("limit"))
		offset, err2 := strconv.Atoi(ctx.Query("offset"))
		if err1 != nil || err2 != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingOffsetLimit.Error(),
			})
			return
		}

		fav, errF := favoriteAPI.GetFavoriteByUserId(ctx.Context(), dec.UserId, limit, offset)
		if errF != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errF.Error(),
			})
			return
		}

		ctx.Status(200).JSON(fav)
	}
}
