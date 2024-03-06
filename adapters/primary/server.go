package primary

import (
	categories "github.com/ecommerce/core/services/category"
	productImages "github.com/ecommerce/core/services/productImage"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

type App struct {
	fiber           *fiber.App
	usersAPI        users.API
	productAPI      products.API
	categoriesAPI   categories.API
	productImageAPI productImages.API
	tokenAPI        ports.TokenService
	storageAPI      ports.StorageService
	emailAPI        ports.EmailService
	port            string
}

func NewApp(tokenAPI ports.TokenService,
	storageAPI ports.StorageService,
	usersAPI users.API,
	productAPI products.API,
	categoryAPI categories.API,
	productImageAPI productImages.API,
	emailAPI ports.EmailService,
	port string) *App {
	newFiber := &App{
		fiber:           fiber.New(),
		usersAPI:        usersAPI,
		productAPI:      productAPI,
		categoriesAPI:   categoryAPI,
		productImageAPI: productImageAPI,
		tokenAPI:        tokenAPI,
		storageAPI:      storageAPI,
		emailAPI:        emailAPI,
		port:            port,
	}
	newFiber.fiber.Use(cors.New())
	initRoutes(newFiber)
	return newFiber
}

func Run(a *App) error {
	return a.fiber.Listen(a.port)
}
