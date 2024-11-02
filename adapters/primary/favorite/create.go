package favorites

import (
	"errors"

	"github.com/ecommerce/core/domain/favorite"
	favorites "github.com/ecommerce/core/services/favorite"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
)

var (
	AuthHeader                 = "Authorization"
	ErrorMissingProductIdField = errors.New("missing product id")
	ErrorAlreadyExistFavorite  = errors.New("product is already added to favorites")
)

func CreateFavorite(favoriteAPI favorites.API) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		dec := ctx.Locals("user").(*ports.Claims)

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
