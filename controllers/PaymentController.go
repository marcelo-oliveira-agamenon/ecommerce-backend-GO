package controller

import (
	"strconv"

	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	b "github.com/ecommerce/utility"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
)

const (
	credit_card  = "CREDIT_CARD"
	money  = "MONEY"
	debit_card  = "DEBIT_CARD"
	transfer  = "TRANSFER"
	pix  = "PIX"
)

//Get all payments by user id
func GetAllPaymentsByUser(w *fiber.Ctx) {
	userid := b.ClaimTokenData(w)
	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))
	var payment []u.Payment

	result := db.DBConn.Model(&u.Payment{}).Limit(limit).Offset(offset).Where("user_id = ?", userid.UserId).Find(&payment)
	if result.Error != nil {
		w.Status(500).JSON("Error listing payments")
		return
	}

	w.Status(200).JSON(payment)
}

//Create new payment by order id
func InsertNewPayment(w *fiber.Ctx)  {
	userid := b.ClaimTokenData(w)
	var payment u.Payment
	payment.OrderID = w.FormValue("order_id")
	payment.UserID = userid.UserId
	payment.TotalValue, _ = strconv.ParseFloat(w.FormValue("total_value"), 64)
	payment.PaidValue, _ = strconv.ParseFloat(w.FormValue("paid_value"), 64)
	paymentMethod := w.FormValue("payment_method")

	if (paymentMethod != credit_card && paymentMethod != money && paymentMethod != debit_card && paymentMethod != transfer && paymentMethod != pix) {
		w.Status(422).JSON("Payment Method not accepted")
		return
	}

	if (payment.OrderID == "" || payment.UserID == "") {
		w.Status(422).JSON("Payment missing user_id or order_id")
		return
	}

	if (w.FormValue("total_value") == "" || w.FormValue("total_value") == "") {
		w.Status(422).JSON("Payment missing total value or paid value")
		return
	}

	if (payment.TotalValue < payment.PaidValue) {
		w.Status(422).JSON("Payment Total value can't be less than paid value")
		return
	}
	
	id, err := uuid.NewV4()
	if err != nil {
		w.Status(500).JSON("Error in uuid generate")
		return
	}
	
	payment.ID = id
	payment.PaymentMethod = paymentMethod
	result := db.DBConn.Create(&payment)
	if result.Error != nil {
		w.Status(500).JSON("Error creating payment")
		return
	}

	w.Status(200).JSON(payment)
}

//Delete payment, soft delete from database
func DeletePayment(w *fiber.Ctx)  {
	paymentId := w.Params("id")
	var payment u.Payment

	result := db.DBConn.Where("id = ?", paymentId).Find(&payment)
	if result.Error != nil {
		w.Status(422).JSON("This payment doesn't exist")
		return
	}

	resultDelete := db.DBConn.Where("id = ?", paymentId).Delete(&payment)
	if resultDelete.Error != nil {
		w.Status(500).JSON("Error delete payment")
		return
	}

	w.Status(200).JSON(payment)
}