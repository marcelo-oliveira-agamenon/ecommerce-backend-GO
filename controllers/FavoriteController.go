package controller

import (
	"strconv"

	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	b "github.com/ecommerce/utility"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/segmentio/ksuid"
)

//CreateFavorite product for a user
func CreateFavorite(w *fiber.Ctx)  {
	userid := b.ClaimTokenData(w)
	if userid.UserId == "" || w.FormValue("productid") == "" {
		w.Status(500).JSON("Missing fields")
		return
	}

	var fav u.Favorites
	fav.ID = ksuid.New().String()
	fav.UserID = uuid.FromStringOrNil(userid.UserId)
	fav.ProductID = w.FormValue("productid")

	result := db.DBConn.Create(&fav)
	if result.Error != nil {
		w.Status(500).JSON("Error creating favorites")
		return
	}

	w.Status(201).JSON(fav)
}

//GetFavoriteByUser params id user
func GetFavoriteByUser(w * fiber.Ctx)  {
	userid := b.ClaimTokenData(w)
	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))

	var fav []u.Favorites
	result := db.DBConn.Where("user_id", userid.UserId).Limit(limit).Offset(offset).Find(&fav)
	if result.Error != nil {
		w.Status(500).JSON("Error set favorite")
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
		w.Status(500).JSON("Error removing favorite")
		return
	}

	w.Status(200).JSON("Sucess removing from favorite")
}