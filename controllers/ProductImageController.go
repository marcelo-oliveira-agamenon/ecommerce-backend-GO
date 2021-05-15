package controller

import (
	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	q "github.com/ecommerce/utility"
	"github.com/gofiber/fiber"
	"github.com/segmentio/ksuid"
)

//Insert image into database, and upload to AWS S3
func InsertProductImage(w *fiber.Ctx) {
	var prodImg u.ProductImage
	productID := w.Params("product_id")

	if productID == "" {
		w.Status(500).JSON("Missing params, product id")
		return
	}

	photo, _ := w.FormFile("img")
	if photo != nil {
		file, err := photo.Open()
		key, url := q.SendImageToAWS(file, photo.Filename, photo.Size, "product")
		if key == "" || err != nil {
			w.Status(500).JSON("Error upload image to AWS")
			return
		}
		defer file.Close()

		prodImg.ImageKey = key
		prodImg.ImageURL = url
	}

	prodImg.ID = ksuid.New().String()
	prodImg.Productid = productID
	result := db.DBConn.Create(&prodImg)
	if result.Error != nil {
		q.DeleteImageInAWS(prodImg.ImageKey)
		w.Status(500).JSON("Error creating product image")
		return
	}

	w.Status(201).JSON(prodImg)
}

//Delete image from database, and from AWS S3
func DeleteProductImage(w* fiber.Ctx) {
	productImageID := w.Params("id")
	var prodImg u.ProductImage

	if productImageID == "" {
		w.Status(500).JSON("Missing params, product id")
		return
	}

	resultProdImg := db.DBConn.Where("id", productImageID).Find(&prodImg)
	if resultProdImg.Error != nil {
		w.Status(422).JSON("This product image doesn't exist")
		return
	}

	result := db.DBConn.Unscoped().Delete(&prodImg)
	if result.Error != nil {
		w.Status(500).JSON("Error deleting product image")
		return
	}

	q.DeleteImageInAWS(prodImg.ImageKey)

	w.Status(200).JSON(prodImg)
}