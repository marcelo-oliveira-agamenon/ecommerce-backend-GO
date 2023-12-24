package primary

import (
	"github.com/ecommerce/adapters/primary/users"
	"github.com/ecommerce/ports"
)

//Get fiber instance and import routes
func initRoutes(a *App, token ports.TokenService) {
	a.fiber.Post("/v1/signUp", users.SignUp(a.usersAPI, token))
}