package routes

import (
	"github.com/gofiber/fiber"

	a "github.com/ecommerce/controllers"
	b "github.com/ecommerce/utility"
)

//Routes create router and import routes
func Routes(router *fiber.App) {
	router.Post("/v1/login", a.Login)
	router.Post("/v1/admin/login", a.LoginAdmin)
	router.Post("/v1/loginWithFacebook", a.LoginWithFacebook)
	router.Post("/v1/signUp", a.SignUpUser)
	router.Patch("/v1/resetPassword", a.ResetPassword)
	router.Post("/v1/resetPasswordLink", a.SendEmailToResetPassword)
	router.Patch("/v1/refresh", a.RefreshToken)
	
	router.Use(b.VerifyToken)
	router.Get("/v1/admin/card1", a.OrdersQuantityByPeriod)
	router.Get("/v1/admin/card2", a.GetProfitOfOrdersByMonths)
	router.Get("/v1/admin/card3", a.GetCountForAdmin)
	router.Get("/v1/admin/card4", a.GetQuantityProductsByCategories)

	router.Patch("/v1/users/toggleRoles/:id", a.ToggleRolesUser)
	router.Get("/v1/product", a.GetAllProducts)
	router.Post("/v1/product", a.InsertProduct)
	router.Get("/v1/product/getbyId/:id", a.GetProductByID)
	router.Get("/v1/product/category/:id", a.GetAllProductsByCategory)
	router.Get("/v1/product/promotion", a.GetProductPromotion)
	router.Get("/v1/product/recent", a.GetRecentProducts)
	router.Get("/v1/product/search/:string", a.GetProductByName)
	
	router.Get("/v1/category", a.SelectCategoryAll)
	router.Post("/v1/category", a.InsertCategory)

	router.Get("/v1/order/user", a.GetByUser)
	router.Post("/v1/order", a.CreateOrder)
	router.Patch("/v1/order/payment/:id/:bool", a.PaymentChangeOrderByID)
	router.Patch("/v1/order/:id/rate/:rate", a.RateOrder)
	router.Patch("/v1/order/:id/status/:status", a.ChangeStatusOrder)

	router.Post("/v1/favorite", a.CreateFavorite)
	router.Get("/v1/favorite", a.GetFavoriteByUser)
	router.Delete("/v1/favorite/:id", a.RemoveFromFavorite)

	router.Post("/v1/product-image/:product_id", a.InsertProductImage)
	router.Delete("/v1/product-image/:id", a.DeleteProductImage)

	router.Get("/v1/payment/user", a.GetAllPaymentsByUser)
	router.Post("/v1/payment", a.InsertNewPayment)
	router.Delete("/v1/payment/:id", a.DeletePayment)

	router.Post("/v1/coupon", a.CreateCoupon)
	router.Get("/v1/coupon", a.VerifyCouponStillActive)
}