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
	
	router.Get("/v1/product", a.GetAllProducts)
	router.Post("/v1/product", a.InsertProduct)
	
	router.Get("/v1/category", a.SelectCategoryAll)
	router.Post("/v1/category", a.InsertCategory)

	router.Post("/v1/order", a.CreateOrder)
	router.Patch("/v1/order/:id", a.PaidOrderByID)
	router.Use(b.VerifyToken)
}