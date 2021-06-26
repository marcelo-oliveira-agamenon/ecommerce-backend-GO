package controller

import (
	"github.com/ecommerce/db"
	a "github.com/ecommerce/models"
	"github.com/gofiber/fiber"
)


type Card3Content struct {
	countAllProducts		int64
	countAllUsers		int64
	countAllOrders		int64
	countAllPaidOrders	int64
}

//Get count of elements for admin, card 3
func GetCountForAdmin(w *fiber.Ctx)  {
	var card3 Card3Content

	resultCountAllProducts := db.DBConn.Table("products").Count(&card3.countAllProducts)
	if resultCountAllProducts.Error != nil {
		w.Status(500).JSON("Error getting count for all products")
		return
	}

	resultCountAllUsers := db.DBConn.Table("users").Count(&card3.countAllUsers)
	if resultCountAllUsers.Error != nil {
		w.Status(500).JSON("Error getting count for all users")
		return
	}

	resultCountAllOrders := db.DBConn.Table("orders").Count(&card3.countAllOrders)
	if resultCountAllOrders.Error != nil {
		w.Status(500).JSON("Error getting count for all orders")
		return
	}

	resultCountAllPaidOrders := db.DBConn.Model(&a.Order{}).Where("paid = ?", "true").Count(&card3.countAllPaidOrders)
	if resultCountAllPaidOrders.Error != nil {
		w.Status(500).JSON("Error getting count for all paid orders")
		return
	}

	w.Status(200).JSON(&fiber.Map{
		"countAllOrders": card3.countAllOrders,
		"countAllPaidOrders": card3.countAllPaidOrders,
		"countAllProducts": card3.countAllProducts,
		"countAllUsers": card3.countAllUsers,
	})
}