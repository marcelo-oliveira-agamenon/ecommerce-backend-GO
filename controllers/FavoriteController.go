package controller

import (
	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/segmentio/ksuid"
)

//CreateFavorite product for a user
func CreateFavorite(w *fiber.Ctx)  {
	if w.FormValue("userid") == "" || w.FormValue("productid") == "" {
		w.Status(500).JSON("Missing fields")
		return
	}

	var fav u.Favorites
	fav.ID = ksuid.New().String()
	fav.UserID = uuid.FromStringOrNil(w.FormValue("userid"))
	fav.ProductID = w.FormValue("productid")

	result := db.DBConn.Create(&fav)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(201).JSON(fav)
}

//GetFavoriteByUser params id user
func GetFavoriteByUser(w * fiber.Ctx)  {
	id := w.Params("id")

	var fav []u.Favorites
	result := db.DBConn.Where("user_id", id).Find(&fav)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(fav)
}

//RemoveFromFavorite params id favorite
func RemoveFromFavorite(w *fiber.Ctx)  {
	id := w.Params("id")

	var fav u.Favorites
	result := db.DBConn.Where("id", id).Delete(&fav)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON("Sucess removing from favorite")
}