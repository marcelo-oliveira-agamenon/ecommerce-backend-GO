package prodImages

import (
	"errors"

	"github.com/ecommerce/core/domain/productImage"
	productImages "github.com/ecommerce/core/services/productImage"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingProductIdField = errors.New("missing product id parameter")
	ErrorMissingImage          = errors.New("missing image")
	ErrorProductIdDoenstExist  = errors.New("product with this id doenst exist")
)

func CreateProductImage(productImageAPI productImages.API, productAPI products.API, storage ports.StorageService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		prodId := ctx.Params("id")
		if prodId == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingProductIdField.Error(),
			})
		}

		_, errPr := productAPI.GetProductById(ctx.Context(), prodId)
		if errPr != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorProductIdDoenstExist.Error(),
			})
		}

		img, errI := ctx.FormFile("img")
		if errI != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorMissingImage.Error(),
			})
		}

		file, err := img.Open()
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		resp, errAdd := storage.SaveFileAWS(file, img.Filename, img.Size, "product")
		if errAdd != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errAdd.Error(),
			})
		}

		prodI, errP := productImageAPI.CreateProductImage(ctx.Context(), productImage.ProductImage{
			Productid: prodId,
			ImageKey:  resp.ImageKey,
			ImageURL:  resp.ImageURL,
		})
		if errP != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errP.Error(),
			})
		}

		return ctx.Status(201).JSON(prodI)
	}
}
