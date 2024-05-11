package primary

import (
	admins "github.com/ecommerce/adapters/primary/admin"
	categories "github.com/ecommerce/adapters/primary/category"
	coupons "github.com/ecommerce/adapters/primary/coupon"
	favorites "github.com/ecommerce/adapters/primary/favorite"
	"github.com/ecommerce/adapters/primary/middleware"
	orders "github.com/ecommerce/adapters/primary/order"
	payments "github.com/ecommerce/adapters/primary/payment"
	prodImages "github.com/ecommerce/adapters/primary/productImage"
	"github.com/ecommerce/adapters/primary/products"
	reports "github.com/ecommerce/adapters/primary/report"
	"github.com/ecommerce/adapters/primary/users"
)

// Get fiber instance and import routes
func initRoutes(a *App) {
	v1 := a.fiber.Group("/v1")
	{
		v1.Post("/login", users.Login(a.usersAPI, a.tokenAPI, a.redisAPI))
		v1.Post("/signUp", users.SignUp(a.usersAPI, a.tokenAPI, a.storageAPI, a.emailAPI, a.kafkaAPI))
		v1.Post("/loginFacebook", users.LoginFacebook(a.usersAPI, a.tokenAPI, a.redisAPI))
		v1.Patch("/resetPassword", users.ResetPassword(a.usersAPI))
		v1.Post("/resetPasswordLink", users.SendEmailResetPassword(a.usersAPI, a.emailAPI))
		v1.Patch("/refresh", users.RefreshToken(a.usersAPI, a.tokenAPI, a.redisAPI))

		authUser := v1.Use(middleware.VerifyToken(a.tokenAPI))
		{
			user := authUser.Group("/users")
			{
				user.Patch("/toggleRoles/:id", users.ToggleRoles(a.usersAPI, a.logAPI))
			}

			product := authUser.Group("/product")
			{
				product.Get("/", products.GetAllProducts(a.productAPI))
				product.Get("/:id", products.GetProductById(a.productAPI))
				product.Post("/", products.CreateProduct(a.productAPI, a.categoriesAPI, a.logAPI, a.tokenAPI))
				product.Put("/:id", products.EditProduct(a.productAPI))
				product.Delete("/:id", products.DeleteProductById(a.productAPI))
			}

			productImage := authUser.Group("/product-image")
			{
				productImage.Post("/:id", prodImages.CreateProductImage(a.productImageAPI, a.productAPI, a.storageAPI))
				productImage.Delete("/:id", prodImages.DeleteProductImage(a.productImageAPI, a.storageAPI))
			}

			category := authUser.Group("/category")
			{
				category.Get("/", categories.GetAllCategories(a.categoriesAPI))
				category.Get("/:id", categories.GetOneCategory(a.categoriesAPI))
				category.Post("/", categories.CreateCategory(a.categoriesAPI, a.storageAPI))
				category.Delete("/:id", categories.DeleteCategory(a.categoriesAPI, a.productAPI))
			}

			favorite := authUser.Group("/favorite")
			{
				favorite.Get("/", favorites.GetFavoriteByUserId(a.favoriteAPI, a.tokenAPI))
				favorite.Post("/", favorites.CreateFavorite(a.favoriteAPI, a.tokenAPI))
				favorite.Delete("/:id", favorites.DeleteFavorite((a.favoriteAPI)))
			}

			coupon := authUser.Group("/coupon")
			{
				coupon.Get("/", coupons.VerifyCouponStillActive(a.couponAPI))
				coupon.Post("/", coupons.CreateCoupon(a.couponAPI))
			}

			order := authUser.Group("/order")
			{
				order.Get("/", orders.GetByUserId(a.orderAPI, a.usersAPI, a.tokenAPI))
				order.Post("/", orders.CreateOrder(a.orderAPI, a.usersAPI, a.productAPI, a.tokenAPI))
				order.Patch("/payment/:id", orders.EditPayment(a.orderAPI))
				order.Patch("/rate/:id", orders.EditRate(a.orderAPI))
				order.Patch("/status/:id", orders.EditStatus(a.orderAPI))
			}

			payment := authUser.Group("/payment")
			{
				payment.Get("/", payments.GetAllByUser(a.paymentAPI, a.usersAPI, a.tokenAPI))
				payment.Post("/", payments.CreatePayment(a.paymentAPI, a.usersAPI, a.orderAPI, a.tokenAPI))
				payment.Delete("/:id", payments.DeletePayment(a.paymentAPI))
			}

			report := authUser.Group("/report")
			{
				report.Get("/users/export", reports.ExportUsers(a.usersAPI))
				report.Post("/users/import", reports.ImportUsers(a.usersAPI))
			}

			admin := authUser.Group("/admin")
			{
				admin.Get("/card1", admins.GetOrdersByPeriod(a.orderAPI))
				admin.Get("/card2", admins.GetProfitByOrdersByMonths(a.orderAPI))
				admin.Get("/card3", admins.GetCountForAdmin(a.usersAPI, a.orderAPI, a.productAPI))
				admin.Get("/card4", admins.GetProductQuantityByCategories(a.productAPI))
			}
		}
	}
}
