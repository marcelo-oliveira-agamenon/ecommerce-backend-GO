package controller

import (
	"strconv"
	"strings"

	"github.com/ecommerce/db"
	e "github.com/ecommerce/models"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/segmentio/ksuid"
)

//GetByUser get orders by userid
func GetByUser(w *fiber.Ctx)  {
	userid := w.Params("id")

	var orders []e.Order
	result := db.DBConn.Where("user_id", userid).Find(&orders)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(orders)
}

//CreateOrder insert a order to database
func CreateOrder(w * fiber.Ctx)  {
	var order e.Order
	if w.FormValue("userID") == "" || w.FormValue("productID") == "" || w.FormValue("qtd") == "" {
		w.Status(500).JSON("Missing fields")
		return
	}
	order.ID = ksuid.New().String()
	order.UserID = uuid.FromStringOrNil(w.FormValue("userID"))
	order.ProductID = w.FormValue("productID")
	order.Qtd, _ = strconv.Atoi(w.FormValue("qtd"))
	order.Paid = false
	order.TotalValue, _ = strconv.ParseFloat(w.FormValue("totalValue"), 64)
	order.Status = "PENDENTE"

	result := db.DBConn.Create(&order)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(201).JSON(order)
}

//PaymentChangeOrderByID its in the name fucker
func PaymentChangeOrderByID(w *fiber.Ctx)  {
	id := w.Params("id")
	status, _ := strconv.ParseBool(w.Params("bool"))
	
	var order e.Order
	order.ID = id
	result := db.DBConn.Model(&order).Update("paid", status)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(order)
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

//ChangeStatusOrder enum status change
func ChangeStatusOrder(w *fiber.Ctx)  {
	id := w.Params("id")
	status := strings.ToUpper(w.Params("status"))
	if status != "PENDENTE" && status != "CANCELADO" && status != "ENTREGUE" && status != "ANDAMENTO" {
		w.Status(500).JSON("Unknown Status")
		return
	}

	var order e.Order
	order.ID = id
	result := db.DBConn.Model(&order).Update("status", status)
	if result.Error != nil {
		w.Status(500).JSON("Server error")
		return
	}

	w.Status(200).JSON(order)
}