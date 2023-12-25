package primary

import (
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

type App struct {
	fiber    *fiber.App
	usersAPI users.API
	tokenAPI ports.TokenService
	storage  ports.StorageService
	port     string
}

func NewApp(tokenAPI ports.TokenService, storageAPI ports.StorageService, usersAPI users.API, port string) *App {
	newFiber := &App{
		fiber:    fiber.New(),
		usersAPI: usersAPI,
		tokenAPI: tokenAPI,
		storage:  storageAPI,
		port:     port,
	}
	newFiber.fiber.Use(cors.New())
	initRoutes(newFiber)
	return newFiber
}

func Run(a *App) error {
	return a.fiber.Listen(a.port)
}
