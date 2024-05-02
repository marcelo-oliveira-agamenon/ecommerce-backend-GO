package primary

import (
	categories "github.com/ecommerce/core/services/category"
	coupons "github.com/ecommerce/core/services/coupon"
	favorites "github.com/ecommerce/core/services/favorite"
	logs "github.com/ecommerce/core/services/log"
	orders "github.com/ecommerce/core/services/order"
	"github.com/ecommerce/core/services/payments"
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
	favoriteAPI     favorites.API
	couponAPI       coupons.API
	orderAPI        orders.API
	paymentAPI      payments.API
	logAPI          logs.API
	tokenAPI        ports.TokenService
	storageAPI      ports.StorageService
	emailAPI        ports.EmailService
	redisAPI        ports.RedisService
	port            string
}

func NewApp(
	tokenAPI ports.TokenService,
	storageAPI ports.StorageService,
	usersAPI users.API,
	productAPI products.API,
	categoryAPI categories.API,
	productImageAPI productImages.API,
	favoriteAPI favorites.API,
	couponAPI coupons.API,
	orderAPI orders.API,
	paymentAPI payments.API,
	logAPI logs.API,
	emailAPI ports.EmailService,
	redisAPI ports.RedisService,
	port string) *App {
	newFiber := &App{
		fiber:           fiber.New(),
		usersAPI:        usersAPI,
		productAPI:      productAPI,
		categoriesAPI:   categoryAPI,
		productImageAPI: productImageAPI,
		favoriteAPI:     favoriteAPI,
		couponAPI:       couponAPI,
		orderAPI:        orderAPI,
		paymentAPI:      paymentAPI,
		logAPI:          logAPI,
		tokenAPI:        tokenAPI,
		storageAPI:      storageAPI,
		emailAPI:        emailAPI,
		redisAPI:        redisAPI,
		port:            port,
	}
	newFiber.fiber.Use(cors.New())
	initRoutes(newFiber)
	return newFiber
}

func Run(a *App) error {
	return a.fiber.Listen(a.port)
}
