package prodImages

import (
	"errors"

	"github.com/ecommerce/core/domain/productImage"
	productImages "github.com/ecommerce/core/services/productImage"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	ErrorMissingProductIdField = errors.New("missing product id parameter")
	ErrorMissingImage          = errors.New("missing image")
	ErrorProductIdDoenstExist  = errors.New("product with this id doenst exist")
)

func CreateProductImage(productImageAPI productImages.API, productAPI products.API, storage ports.StorageService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		prodId := ctx.Params("id")
		if prodId == "" {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingProductIdField.Error(),
			})
			return
		}

		_, errPr := productAPI.GetProductById(ctx.Context(), prodId)
		if errPr != nil {
			ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorProductIdDoenstExist.Error(),
			})
			return
		}

		img, errI := ctx.FormFile("img")
		if errI != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorMissingImage.Error(),
			})
			return
		}

		file, err := img.Open()
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}
		resp, errAdd := storage.SaveFileAWS(file, img.Filename, img.Size, "product")
		if errAdd != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errAdd.Error(),
			})
			return
		}

		prodI, errP := productImageAPI.CreateProductImage(ctx.Context(), productImage.ProductImage{
			Productid: prodId,
			ImageKey:  resp.ImageKey,
			ImageURL:  resp.ImageURL,
		})
		if errP != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
			return
		}

		ctx.Status(201).JSON(prodI)
	}
}
