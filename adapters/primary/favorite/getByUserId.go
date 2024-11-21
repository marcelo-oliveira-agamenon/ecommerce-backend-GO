package favorites

import (
	"errors"
	"strconv"

	favorites "github.com/ecommerce/core/services/favorite"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingOffsetLimit = errors.New("missing limit or offset query parameter")
)

func GetFavoriteByUserId(favoriteAPI favorites.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dec := ctx.Locals("user").(*ports.Claims)

		limit, err1 := strconv.Atoi(ctx.Query("limit"))
		offset, err2 := strconv.Atoi(ctx.Query("offset"))
		if err1 != nil || err2 != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingOffsetLimit.Error(),
			})
		}

		fav, errF := favoriteAPI.GetFavoriteByUserId(ctx.Context(), dec.UserId, limit, offset)
		if errF != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errF.Error(),
			})
		}

		return ctx.Status(200).JSON(fav)
	}
}
