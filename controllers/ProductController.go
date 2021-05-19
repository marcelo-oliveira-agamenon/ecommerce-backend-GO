package controller

import (
	"strconv"

	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
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
	aux.Rate = 0

	result := db.DBConn.Create(&aux)
	if result.Error != nil {
		w.Status(500).JSON("Error creating product")
		return
	}

	w.Status(201).JSON(aux)
}

//GetAllProducts from database
func GetAllProducts(w *fiber.Ctx) {
	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))

	var products []u.Product

	result := db.DBConn.Preload("ProductImage").Joins("Category").Limit(limit).Offset(offset).Find(&products)
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
	result := db.DBConn.Preload("ProductImage").Joins("Category").Where("products.id", productID).Find(&product)
	if result.Error != nil {
		w.Status(500).JSON("Error listing product")
		return
	}

	w.Status(200).JSON(product)
}

//GetAllProductsByCategory from database
func GetAllProductsByCategory(w *fiber.Ctx)  {
	categoryID := w.Params("id")
	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))

	var products []u.Product
	result := db.DBConn.Preload("ProductImage").Joins("Category").Where("products.categoryid", categoryID).Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}

//GetProductPromotion from database
func GetProductPromotion(w *fiber.Ctx)  {
	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))

	var products []u.Product
	result := db.DBConn.Preload("ProductImage").Joins("Category").Where("has_promotion", true).Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}

//GetRecentProducts from database
func GetRecentProducts(w *fiber.Ctx)  {
	var products []u.Product
	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))

	result := db.DBConn.Preload("ProductImage").Joins("Category").Order("created_at desc").Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}

//GetProductByName from database
func GetProductByName(w *fiber.Ctx)  {
	strv := w.Params("string")
	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))

	var products []u.Product
	result := db.DBConn.Preload("ProductImage").Joins("Category").Where("products.name like ?", "%" + strv + "%").Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	w.Status(200).JSON(products)
}