package controller

import (
	"github.com/ecommerce/db"
	e "github.com/ecommerce/models"
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
		w.Status(404).JSON("Missing field name")
		return
	}

	var category e.Category
	category.Name = w.FormValue("name")
	result := db.DBConn.Create(&category)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(category)
}