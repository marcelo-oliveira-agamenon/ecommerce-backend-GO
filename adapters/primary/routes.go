package primary

import (
	"github.com/ecommerce/adapters/primary/users"
)

// Get fiber instance and import routes
func initRoutes(a *App) {
	a.fiber.Post("/v1/login", users.Login(a.usersAPI, a.tokenAPI))
	a.fiber.Post("/v1/signUp", users.SignUp(a.usersAPI, a.tokenAPI, a.storage))
}
