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

	var prImage  u.ProductImage
	photos, _ := w.FormFile("photos")
	if photos != nil {
		file, err := photos.Open()
		key, url := q.SendImageToAWS(file, photos.Filename, photos.Size, "product")
		if key == "" || err != nil {
			w.Status(500).JSON("Error upload image to AWS")
			return
		}
		defer file.Close()

		prImage.ImageKey = key
		prImage.ImageURL = url
	}

	result := db.DBConn.Create(&aux)
	if result.Error != nil {
		w.Status(500).JSON("Error creating product")
		q.DeleteImageInAWS(prImage.ImageKey)
		return
	}

	prImage.ID = ksuid.New().String()
	prImage.Productid = aux.ID
	resultImage := db.DBConn.Create(&prImage)
	if resultImage.Error != nil {
		w.Status(500).JSON("Error creating product image")
		return
	}

	w.Status(201).JSON(aux)
}

//GetAllProducts from database
func GetAllProducts(w *fiber.Ctx) {
	var products []u.Product
	result := db.DBConn.Preload("ProductImage").Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}

//GetProductByID from database
func GetProductByID(w *fiber.Ctx) {
	productID := w.Params("id")

	var product u.Product
	result := db.DBConn.Preload("ProductImage").Where("id", productID).Find(&product)
	if result.Error != nil {
		w.Status(500).JSON("Error listing product")
		return
	}

	w.Status(200).JSON(product)
}

//GetAllProductsByCategory from database
func GetAllProductsByCategory(w *fiber.Ctx)  {
	categoryID := w.Params("id")

	var products []u.Product
	result := db.DBConn.Preload("ProductImage").Where("categoryid", categoryID).Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}

//GetProductPromotion from database
func GetProductPromotion(w *fiber.Ctx)  {
	var products []u.Product
	result := db.DBConn.Preload("ProductImage").Where("has_promotion", true).Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}

//GetRecentProducts from database
func GetRecentProducts(w *fiber.Ctx)  {
	var products []u.Product

	result := db.DBConn.Preload("ProductImage").Order("created_at desc").Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}

//GetProductByName from database
func GetProductByName(w *fiber.Ctx)  {
	strv := w.Params("string")

	var products []u.Product
	result := db.DBConn.Preload("ProductImage").Where("name like ?", "%" + strv + "%").Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}