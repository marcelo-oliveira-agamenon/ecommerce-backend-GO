package controller

import (
	"strconv"

	"github.com/ecommerce/db"
	e "github.com/ecommerce/models"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/segmentio/ksuid"
)

//CreateOrder insert a order to database
func CreateOrder(w * fiber.Ctx)  {
	var order e.Order
	if w.FormValue("userID") == "" || w.FormValue("productID") == "" || w.FormValue("qtd") == "" {
		w.Status(403).JSON("Missing fields")
		return
	}
	order.ID = ksuid.New().String()
	order.UserID = uuid.FromStringOrNil(w.FormValue("userID"))
	order.ProductID = w.FormValue("productID")
	order.Qtd, _ = strconv.Atoi(w.FormValue("qtd"))
	order.Paid = false

	result := db.DBConn.Create(&order)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(201).JSON(order)
}

//PaidOrderByID its in the name fucker
func PaidOrderByID(w *fiber.Ctx)  {
	id := w.Params("id")
	
	var order e.Order
	order.ID = id
	result := db.DBConn.Model(&order).Update("paid", true)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(&fiber.Map{
		"order": order,
		"message": "Order paid",
	})
}

//CancelPayOrderByID its in the name fucker
func CancelPayOrderByID(w *fiber.Ctx)  {
	var order e.Order
	order.ID = w.Params("id")
	result := db.DBConn.Model(&order).Update("paid", false)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(&fiber.Map{
		"order": order,
		"message": "Order payment cancelled",
	})
}

//RateOrder put rate service
func RateOrder(w *fiber.Ctx)  {
	var order e.Order
	order.ID = w.Params("id")
	order.Rate, _ = strconv.Atoi(w.Params("rate"))

	result := db.DBConn.Model(&order).Update("rate", order.Rate)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON("Order rated")
}