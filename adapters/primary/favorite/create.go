package favorites

import (
	"errors"

	"github.com/ecommerce/core/domain/favorite"
	favorites "github.com/ecommerce/core/services/favorite"
	"github.com/ecommerce/core/util"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
)

var (
	AuthHeader                 = "Authorization"
	ErrorMissingProductIdField = errors.New("missing product id")
	ErrorAlreadyExistFavorite  = errors.New("product is already added to favorites")
)

func CreateFavorite(favoriteAPI favorites.API, token ports.TokenService) fiber.Handler {
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

		prodId := ctx.FormValue("productid")
		if prodId == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingProductIdField.Error(),
			})
			return
		}

		favs, errG := favoriteAPI.GetFavoriteByUserIdAndProductId(ctx.Context(), dec.UserId, prodId)
		if errG != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
			return
		}
		if len(*favs) > 0 {
			ctx.Status(409).JSON(&fiber.Map{
				"error": ErrorAlreadyExistFavorite.Error(),
			})
			return
		}

		fav, errF := favoriteAPI.AddFavorite(ctx.Context(), favorite.Favorite{
			UserID:    uuid.FromStringOrNil(dec.UserId),
			ProductID: prodId,
		})
		if errF != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errF.Error(),
			})
			return
		}

		ctx.Status(201).JSON(fav)
	}
}
