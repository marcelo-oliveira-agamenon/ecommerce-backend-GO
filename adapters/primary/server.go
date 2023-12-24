package primary

import (
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

type App struct {
	fiber		*fiber.App
	usersAPI	users.API
	port		string
}

func NewApp(token ports.TokenService, usersAPI users.API, port string) *App {
	newFiber := &App{
		fiber: fiber.New(),
		usersAPI: usersAPI,
		port: port,
	}
	newFiber.fiber.Use(cors.New())
	initRoutes(newFiber, token)
	return newFiber
}

func Run(a *App) error  {
	return a.fiber.Listen(a.port)
}