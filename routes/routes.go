package routes

import (
	"github.com/gofiber/fiber"

	a "github.com/ecommerce/controllers"
	b "github.com/ecommerce/utility"
)

//Routes create router and import routes
func Routes(router *fiber.App) {
	router.Post("/v1/login", a.Login)
	router.Post("/v1/signUp", a.SignUpUser)
	
	router.Use(b.VerifyToken)
	router.Get("/v1/product", a.GetAllProducts)
	router.Post("/v1/product", a.InsertProduct)
	router.Get("/v1/product/category/:id", a.GetAllProductsByCategory)
	router.Get("/v1/product/promotion", a.GetProductPromotion)
	router.Get("/v1/product/recent", a.GetRecentProducts)
	router.Get("/v1/product/search/:string", a.GetProductByName)
	
	router.Get("/v1/category", a.SelectCategoryAll)
	router.Post("/v1/category", a.InsertCategory)

	router.Post("/v1/order", a.CreateOrder)
	router.Patch("/v1/order/pay/:id", a.PaidOrderByID)
	router.Patch("/v1/order/cancelPay/:id", a.CancelPayOrderByID)
	router.Patch("/v1/order/:id/rate/:rate", a.RateOrder)

	router.Post("/v1/favorite", a.CreateFavorite)
	router.Get("/v1/favorite/:id", a.GetFavoriteByUser)
	router.Delete("/v1/favorite/:id", a.RemoveFromFavorite)
}