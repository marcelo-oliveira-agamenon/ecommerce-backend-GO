package controller

import (
	"strconv"

	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	q "github.com/ecommerce/utility"
	"github.com/gofiber/fiber"
	"github.com/segmentio/ksuid"
)

//InsertProduct to database
func InsertProduct(w *fiber.Ctx) {
	var aux u.Product
	aux.ID = ksuid.New().String()
	aux.Name = w.FormValue("name")
	aux.Categoryid = w.FormValue("categoryid")
	aux.Value, _ = strconv.ParseFloat(w.FormValue("value"), 64)
	aux.StockQtd, _ = strconv.Atoi(w.FormValue("stockqtd"))
	aux.Description = w.FormValue("description")
	aux.TypeUnit = w.FormValue("type")
	aux.TecnicalDetails = w.FormValue("tecnicalDetails")
	aux.HasPromotion, _ = strconv.ParseBool(w.FormValue("hasPromotion"))
	aux.Discount, _ = strconv.ParseFloat(w.FormValue("discount"), 64)
	aux.HasShipping, _ = strconv.ParseBool(w.FormValue("hasShipping"))
	aux.ShippingPrice, _ = strconv.ParseFloat(w.FormValue("shippingPrice"), 64)

	var keyAux string
	photos, _ := w.FormFile("photos")
	if photos != nil {
		file, err := photos.Open()
		photosResponse, key := q.SendImageToAWS(file, photos.Filename, photos.Size, "product")
		if photosResponse == nil || err != nil {
			w.Status(500).JSON("Error upload image to AWS")
			return
		}
		defer file.Close()

		aux.Photos = photosResponse
		keyAux = key
	} else {
		aux.Photos = u.JSONB{}
	}

	result := db.DBConn.Create(&aux)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		q.DeleteImageInAWS(keyAux)
		return
	}

	w.Status(201).JSON(aux)
}

//GetAllProducts from database
func GetAllProducts(w *fiber.Ctx) {
	var products []u.Product
	result := db.DBConn.Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(products)
}

//GetProductByID from database
func GetProductByID(w *fiber.Ctx) {
	productID := w.Params("id")

	var product u.Product
	result := db.DBConn.Where("id", productID).Find(&product)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(product)
}

//GetAllProductsByCategory from database
func GetAllProductsByCategory(w *fiber.Ctx)  {
	categoryID := w.Params("id")

	var products []u.Product
	result := db.DBConn.Where("categoryid", categoryID).Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(products)
}

//GetProductPromotion from database
func GetProductPromotion(w *fiber.Ctx)  {
	var products []u.Product
	result := db.DBConn.Where("has_promotion", true).Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(products)
}

//GetRecentProducts from database
func GetRecentProducts(w *fiber.Ctx)  {
	var products []u.Product

	result := db.DBConn.Order("created_at desc").Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(products)
}

//GetProductByName from database
func GetProductByName(w *fiber.Ctx)  {
	strv := w.Params("string")

	var products []u.Product
	result := db.DBConn.Where("name like ?", "%" + strv + "%").Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(products)
}