package controller

import (
	"strconv"
	"time"

	"github.com/ecommerce/db"
	"github.com/ecommerce/models"
	"github.com/gofiber/fiber"
	"github.com/segmentio/ksuid"
)

//Insert coupon in database
func CreateCoupon(w *fiber.Ctx) {
	var coupon models.Coupon
	var verifyCoupon []models.Coupon

	if (w.FormValue("vality_date") == "" || w.FormValue("discount") == "" || w.FormValue("title") == "") {
		w.Status(422).JSON("Missing fields, vality date or discount")
		return
	}

	resultSearch := db.DBConn.Where("hash = ?", w.FormValue("title")).Find(&verifyCoupon)
	if resultSearch.Error != nil {
		w.Status(500).JSON("Error finding coupon hash")
		return
	}

	if len(verifyCoupon) > 0 {
		w.Status(403).JSON("Coupon hash already exist in database")
		return
	}

	coupon.Discount, _ = strconv.Atoi(w.FormValue("discount"))
	coupon.ValityDate, _ = time.Parse("2006-01-02T15:04:05", w.FormValue("vality_date"))
	coupon.ID = ksuid.New().String()
	coupon.Hash = w.FormValue("title")

	result := db.DBConn.Create(&coupon)
	if result.Error != nil {
		w.Status(500).JSON("Error creating coupon")
		return
	}

	w.Status(201).JSON(coupon)
}

//Verify coupon date vality
func VerifyCouponStillActive(w *fiber.Ctx) {
	var coupon models.Coupon

	hash := w.Query("hash")

	result := db.DBConn.Where("hash = ?", hash).Find(&coupon)
	if result.Error != nil {
		w.Status(500).JSON("Error: there is no coupon with hash informed")
		return
	}

	if coupon.ValityDate.Before(time.Now()) {
		w.Status(200).JSON(&fiber.Map{
			"valid": false,
		})		
		return
	}

	w.Status(200).JSON(&fiber.Map{
		"valid": true,
		"discount": coupon.Discount,
	})
}