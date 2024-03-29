package primary

import (
	categories "github.com/ecommerce/adapters/primary/category"
	favorites "github.com/ecommerce/adapters/primary/favorite"
	"github.com/ecommerce/adapters/primary/middleware"
	prodImages "github.com/ecommerce/adapters/primary/productImage"
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
				product.Get("/:id", products.GetProductById(a.productAPI))
				product.Post("/", products.CreateProduct(a.productAPI, a.categoriesAPI, a.tokenAPI))
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
				favorite.Post("/", favorites.CreateFavorite(a.favoriteAPI, a.tokenAPI))
			}
		}
	}
}
