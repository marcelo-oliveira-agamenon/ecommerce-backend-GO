package favorites

import (
	"errors"
	"strconv"

	favorites "github.com/ecommerce/core/services/favorite"
	"github.com/ecommerce/core/util"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingOffsetLimit = errors.New("missing limit or offset query parameter")
)

func GetFavoriteByUserId(favoriteAPI favorites.API, token ports.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		tok, errT := util.GetToken(ctx, AuthHeader)
		if errT != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": errT.Error(),
			})
			return
		}

		dec, errC := token.ClaimTokenData(*tok)
		if errC != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"error": errC.Error(),
			})
			return
		}

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
