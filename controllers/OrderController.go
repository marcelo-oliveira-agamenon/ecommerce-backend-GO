package controller

import (
	"strconv"
	"strings"
	"time"

	"github.com/ecommerce/db"
	e "github.com/ecommerce/models"
	u "github.com/ecommerce/utility"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type APIOrder struct {
	gorm.Model
	ID				string
	User_id			uuid.UUID
	ProductID			pq.StringArray		`gorm:"type:varchar(64)[]"`
	TotalValue			float64
	Status			string
	Qtd				int
	Paid				bool
	Rate				int
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			gorm.DeletedAt
}

type TemplateDataCreateOrder struct {
	Name			string
	Year		 	string
	OrderNumber		string
	ProductName		string
	OrderValue		string
	OrderQtd		int
}

//GetByUser get orders by userid
func GetByUser(w *fiber.Ctx)  {
	userid := u.ClaimTokenData(w)

	limit, _ := strconv.Atoi(w.Query("limit"))
	offset, _ := strconv.Atoi(w.Query("offset"))

	var orders []APIOrder
	result := db.DBConn.Model(&e.Order{}).Where("user_id", userid.UserId).Limit(limit).Offset(offset).Order("created_at desc").Find(&orders)
	if result.Error != nil {
		w.Status(500).JSON("Error listing orders")
		return
	}

	w.Status(200).JSON(orders)
}

//CreateOrder insert a order to database
func CreateOrder(w *fiber.Ctx)  {
	userid := u.ClaimTokenData(w)
	var order e.Order
	if userid.UserId == "" || len(w.FormValue("productID")) == 0 || w.FormValue("qtd") == "" {
		w.Status(500).JSON("Missing fields")
		return
	}

	order.ID = ksuid.New().String()
	order.Userid = uuid.FromStringOrNil(userid.UserId)
	aux := strings.Split(w.FormValue("productID"), ",")
	order.ProductID = aux
	order.Qtd, _ = strconv.Atoi(w.FormValue("qtd"))
	order.Paid = false
	order.TotalValue, _ = strconv.ParseFloat(w.FormValue("totalValue"), 64)
	order.Status = "PENDENTE"

	products := make([]e.Product, len(order.ProductID))
	for _, s := range order.ProductID {
		var product e.Product
		result:= db.DBConn.Where("id = ?", s).Find(&product)
		if result.Error != nil {
			w.Status(500).JSON("Error finding product for email")
			return
		}

		products = append(products, product)
	}

	if len(products) == 0 {
		w.Status(500).JSON("Error, no valid products")
		return
	}

	result := db.DBConn.Create(&order)
	if result.Error != nil {
		w.Status(500).JSON("Error creating order")
		return
	}

	var user e.User
	result1 := db.DBConn.Where("id = ?", userid.UserId).Find(&user)
	if result1.Error != nil {
		w.Status(500).JSON("Error finding user")
		return
	}

	body := TemplateDataCreateOrder{
		Name: user.Name,
		Year: strconv.Itoa(time.Now().Year()),
		OrderNumber: order.ID,
		ProductName: "dasd",
		OrderValue: w.FormValue("totalValue"),
		OrderQtd: order.Qtd,
	}
	
	u.SendEmailUtility(user.Email, "template/newOrder.html", body, "Seu Pedido #" + order.ID + " foi realizado na Grab and Cash")

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
		w.Status(500).JSON("Error change status")
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
		w.Status(500).JSON("Error set rate")
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
		w.Status(500).JSON("Error set status")
		return
	}

	w.Status(200).JSON(order)
}