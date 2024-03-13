package prodImages

import (
	"errors"

	productImages "github.com/ecommerce/core/services/productImage"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	ErrorIsNotDeleted = errors.New("removing product image")
)

func DeleteProductImage(productImageAPI productImages.API, storage ports.StorageService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		prodId := ctx.Params("id")
		if prodId == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingProductIdField.Error(),
			})
			return
		}

		prodIm, errG := productImageAPI.GetProductImageById(ctx.Context(), prodId)
		if errG != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
			return
		}

		_, errAdd := storage.DeleteFileAWS(prodIm.ImageKey)
		if errAdd != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errAdd.Error(),
			})
			return
		}

		isDel, errP := productImageAPI.DeleteProductImage(ctx.Context(), *prodIm)
		if errP != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
			return
		}

		if isDel {
			ctx.Status(200).JSON(prodIm)
			return
		}

		ctx.Status(500).JSON(&fiber.Map{
			"error": ErrorIsNotDeleted.Error(),
		})
	}
}
