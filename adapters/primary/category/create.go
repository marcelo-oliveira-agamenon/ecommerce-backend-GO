package categories

import (
	"errors"

	"github.com/ecommerce/core/domain/category"
	categories "github.com/ecommerce/core/services/category"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorMissingNameField = errors.New("missing category name")
	ErrorMissingImage     = errors.New("missing category image")
)

func CreateCategory(categoryAPI categories.API, storage ports.StorageService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		name := ctx.FormValue("name")
		if name == "" {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorMissingNameField.Error(),
			})
		}

		img, errI := ctx.FormFile("image")
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
		resp, errAdd := storage.SaveFileAWS(file, img.Filename, img.Size, "category")
		if errAdd != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errAdd.Error(),
			})
		}

		cat, errA := categoryAPI.AddCategory(ctx.Context(), category.Category{
			Name:     name,
			ImageKey: resp.ImageKey,
			ImageURL: resp.ImageURL,
		})
		if errA != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errA.Error(),
			})
		}

		return ctx.Status(201).JSON(cat)
	}
}
