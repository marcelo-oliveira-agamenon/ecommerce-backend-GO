package controller

import (
	"strconv"

	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	"github.com/ecommerce/utility"
	"github.com/gofiber/fiber"
	"github.com/segmentio/ksuid"
)

//InsertProduct to database
func InsertProduct(w *fiber.Ctx) {
	userid := utility.ClaimTokenData(w)
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

	utility.InsertLogRegistryIntoDabatase("product", "Insert product with ID " + aux.ID, userid.UserId)

	w.Status(201).JSON(aux)
}

//EditProduct from database
func EditProductById(w *fiber.Ctx)  {
	productId := w.Params("id")
	product := new(u.Product)

	result := db.DBConn.Where("products.id", productId)
	if result.Error != nil {
		w.Status(404).JSON("This product Id doenst exist on our database")
		return
	}

	if err := w.BodyParser(product); err != nil {
		w.Status(422).JSON("Error: Missing fields on Product")
		return
	}

	resultUpdate := db.DBConn.Save(&product)
	if resultUpdate.Error != nil {
		w.Status(500).JSON("Error editing product")
		return
	}

	w.Status(200).JSON(product)
}

//DeleteProduct from database
func DeleteProductById(w *fiber.Ctx)  {
	productID := w.Params("id")
	var product u.Product

	resultProduct := db.DBConn.Where("products.id", productID).Find(&product)
	if resultProduct.Error != nil {
		w.Status(404).JSON("This product Id doenst exist on our database")
		return
	}

	resultDelete := db.DBConn.Unscoped().Delete(&product)
	if resultDelete.Error != nil {
		w.Status(404).JSON("Error deleting this product")
		return
	}

	w.Status(200).JSON("Product deleted")
}

//GetAllProducts from database
func GetAllProducts(w *fiber.Ctx) {
	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))

	//query params
	getByCategory := w.Query("category")
	getByPromotion := w.Query("promotion")
	getRecentOnes := w.Query("recent")
	getByName := w.Query("name")

	var products []u.Product

	baseQuery := db.DBConn.Preload("ProductImage").Joins("Category").Limit(limit).Offset(offset)

	if getByCategory != "" {
		baseQuery = baseQuery.Where("products.categoryid", getByCategory)
	}

	if getByPromotion != "" {
		baseQuery = baseQuery.Where("has_promotion", true)
	}

	if getRecentOnes != "" {
		baseQuery = baseQuery.Order("created_at desc")
	}

	if getByName != "" {
		baseQuery = baseQuery.Where("products.name like ?", "%" + getByName + "%")
	}

	result := baseQuery.Find(&products)
	if result.Error != nil {
		w.Status(500).JSON("Error listing products")
		return
	}

	var countProducts int64
	countResult := db.DBConn.Table("products").Count(&countProducts)
	if countResult.Error != nil {
		w.Status(500).JSON("Error counting products")
		return
	}

	w.Status(200).JSON(&fiber.Map{
		"products": products,
		"count": countProducts,
	})
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