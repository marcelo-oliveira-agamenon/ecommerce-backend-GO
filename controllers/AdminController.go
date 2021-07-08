package controller

import (
	"time"

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

type AuxTotal struct {
	OrderId		string
	Subtotal		float64
	Month			string
}

type Card2Content struct {
	Month			string
	Data			[]AuxTotal
}

type Card1Content struct {
	Month			string
	Quantity		int64
}

type CountCategories struct {
	Name			string
	Count			int64
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

//Get Orders quantity by month
func OrdersQuantityByPeriod(w *fiber.Ctx)  {
	var localOrders []a.Order
	var totalOrders []Card1Content
	months := make([]string, 5)

	initDatePeriod := time.Now()
	finishDatePeriod := initDatePeriod.AddDate(0, -5, 0)

	result := db.DBConn.Model(&a.Order{}).Where("created_at BETWEEN ? AND ?", finishDatePeriod, initDatePeriod).Find(&localOrders)
	if result.Error != nil {
		w.Status(500).JSON("Error in get orders")
		return
	}

	for i := 0; i < 5; i++ {
		aux := initDatePeriod.AddDate(0, -i, 0)

		months[i] = aux.Month().String()
	}

	for i := 0; i < len(months); i++ {
		var auxOrder Card1Content
		var count int64
		auxOrder.Month = months[i]

		for j := 0; j < len(localOrders); j++ {
			if localOrders[j].CreatedAt.Month().String() == months[i] {
				count++
			}
		}

		auxOrder.Quantity = count
		totalOrders = append(totalOrders, auxOrder)
	}

	w.Status(200).JSON(totalOrders)
}

//Get profit by month, of orders
func GetProfitOfOrdersByMonths(w *fiber.Ctx)  {
	var localOrders []a.Order
	var formatTotal []AuxTotal
	var auxTotal []Card2Content
	months := make([]string, 5)

	initDatePeriod := time.Now()
	finishDatePeriod := initDatePeriod.AddDate(0, -5, 0)

	for i := 0; i < 5; i++ {
		aux := initDatePeriod.AddDate(0, -i, 0)

		months[i] = aux.Month().String()
	}

	result := db.DBConn.Model(&a.Order{}).Where("created_at BETWEEN ? AND ?", finishDatePeriod, initDatePeriod).Find(&localOrders)
	if result.Error != nil {
		w.Status(500).JSON("Error in get orders")
		return
	}

	for i := 0; i < len(localOrders); i++ {
		var aux AuxTotal
		aux.Subtotal = localOrders[i].TotalValue * float64(localOrders[i].Qtd)
		aux.OrderId = localOrders[i].ID
		aux.Month = localOrders[i].CreatedAt.Month().String()

		formatTotal = append(formatTotal, aux)
	}

	for i := 0; i < len(months); i++ {
		var auxMonth Card2Content
		auxMonth.Month = months[i]

		for j := 0; j < len(formatTotal); j++ {
			if formatTotal[j].Month == months[i] {
				auxMonth.Data = append(auxMonth.Data, formatTotal[j])
			}
		}

		auxTotal = append(auxTotal, auxMonth)
	}

	w.Status(200).JSON(auxTotal)
}

//Get quantity of products by categories
func GetQuantityProductsByCategories(w *fiber.Ctx)  {
	var count []CountCategories
	var total int64

	resultTotal := db.DBConn.Raw("select c.name, count(*) from products p join categories c on c.id = p.categoryid group by c.name").Scan(&count)
	if resultTotal.Error != nil {
		w.Status(500).JSON("Error in count products")
		return
	}

	for i := 0; i < len(count); i++ {
		total = total + count[i].Count
	}

	w.Status(200).JSON(&fiber.Map{
		"data": count,
		"total": total,
	})
}