package prodImages

import (
	"errors"

	productImages "github.com/ecommerce/core/services/productImage"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorIsNotDeleted = errors.New("removing product image")
)

func DeleteProductImage(productImageAPI productImages.API, storage ports.StorageService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		prodId := ctx.Params("id")
		if prodId == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingProductIdField.Error(),
			})
		}

		prodIm, errG := productImageAPI.GetProductImageById(ctx.Context(), prodId)
		if errG != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errG.Error(),
			})
		}

		_, errAdd := storage.DeleteFileAWS(prodIm.ImageKey)
		if errAdd != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errAdd.Error(),
			})
		}

		isDel, errP := productImageAPI.DeleteProductImage(ctx.Context(), *prodIm)
		if errP != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
		}

		if isDel {
			return ctx.Status(200).JSON(prodIm)
		}

		return ctx.Status(500).JSON(&fiber.Map{
			"error": ErrorIsNotDeleted.Error(),
		})
	}
}
