package primary

import (
	"github.com/ecommerce/adapters/primary/middleware"
	"github.com/ecommerce/adapters/primary/products"
	"github.com/ecommerce/adapters/primary/users"
)

// Get fiber instance and import routes
func initRoutes(a *App) {
	v1 := a.fiber.Group("/v1")
	{
		v1.Post("/login", users.Login(a.usersAPI, a.tokenAPI))
		v1.Post("/signUp", users.SignUp(a.usersAPI, a.tokenAPI, a.storageAPI))
		v1.Post("/loginFacebook", users.LoginFacebook(a.usersAPI, a.tokenAPI))
		v1.Patch("/resetPassword", users.ResetPassword(a.usersAPI))
		v1.Post("/resetPasswordLink", users.SendEmailResetPassword(a.usersAPI, a.emailAPI))

		authUser := v1.Use(middleware.VerifyToken(a.tokenAPI))
		{
			user := authUser.Group("/users")
			{
				user.Patch("/toggleRoles/:id", users.ToggleRoles(a.usersAPI))
			}

			product := authUser.Group("/product")
			{
				product.Get("/", products.GetAllProducts(a.productAPI))
				product.Post("/", products.CreateProduct(a.productAPI, a.categoriesAPI, a.tokenAPI))
				product.Put("/:id", products.EditProduct(a.productAPI))
				product.Delete("/:id", products.DeleteProductById(a.productAPI))
			}
		}
	}
}
