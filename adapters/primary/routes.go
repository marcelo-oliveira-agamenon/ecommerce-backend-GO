package primary

import (
	"github.com/ecommerce/adapters/primary/users"
)

// Get fiber instance and import routes
func initRoutes(a *App) {
	v1 := a.fiber.Group("/v1")
	{
		v1.Post("/login", users.Login(a.usersAPI, a.tokenAPI))
		v1.Post("/signUp", users.SignUp(a.usersAPI, a.tokenAPI, a.storage))
		v1.Post("/loginFacebook", users.LoginFacebook(a.usersAPI, a.tokenAPI))
		v1.Patch("/resetPassword", users.ResetPassword(a.usersAPI))
	}
}
