package controller

import (
	"github.com/ecommerce/db"
	e "github.com/ecommerce/models"
	q "github.com/ecommerce/utility"
	"github.com/gofiber/fiber"
)

//SelectCategoryAll from database
func SelectCategoryAll(w *fiber.Ctx)  {
	var category []e.Category
	result := db.DBConn.Find(&category)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}
	w.Status(200).JSON(category)
}

//InsertCategory into database
func InsertCategory(w *fiber.Ctx)  {
	if w.FormValue("name") == "" {
		w.Status(500).JSON("Missing field name")
		return
	}

	image, err := w.FormFile("image")
	if err != nil {
		w.Status(500).JSON("Missing field image: " + err.Error())
		return
	}

	file, err := image.Open()
	key, url := q.SendImageToAWS(file, image.Filename, image.Size, "category")
	if key == "" || err != nil {
		w.Status(500).JSON("Error upload image")
		return
	}
	defer file.Close()

	var category e.Category
	category.Name = w.FormValue("name")
	category.ImageKey = key
	category.ImageURL = url
	
	result := db.DBConn.Create(&category)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		q.DeleteImageInAWS(key)
		return
	}

	w.Status(200).JSON(category)
}