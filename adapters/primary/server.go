package primary

import (
	categories "github.com/ecommerce/core/services/category"
	coupons "github.com/ecommerce/core/services/coupon"
	favorites "github.com/ecommerce/core/services/favorite"
	logs "github.com/ecommerce/core/services/log"
	"github.com/ecommerce/core/services/misc"
	orders "github.com/ecommerce/core/services/order"
	ordersdetails "github.com/ecommerce/core/services/ordersDetails"
	"github.com/ecommerce/core/services/payments"
	productImages "github.com/ecommerce/core/services/productImage"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	orderDetailsAPI ordersdetails.API
	paymentAPI      payments.API
	logAPI          logs.API
	miscAPI         misc.API
	tokenAPI        ports.TokenService
	storageAPI      ports.StorageService
	redisAPI        ports.RedisService
	kafkaAPI        ports.KafkaService
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
	orderDetailsAPI ordersdetails.API,
	paymentAPI payments.API,
	logAPI logs.API,
	miscAPI misc.API,
	redisAPI ports.RedisService,
	kafkaAPI ports.KafkaService,
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
		orderDetailsAPI: orderDetailsAPI,
		paymentAPI:      paymentAPI,
		logAPI:          logAPI,
		miscAPI:         miscAPI,
		tokenAPI:        tokenAPI,
		storageAPI:      storageAPI,
		redisAPI:        redisAPI,
		kafkaAPI:        kafkaAPI,
		port:            port,
	}
	newFiber.fiber.Use(cors.New())
	initRoutes(newFiber)

	return newFiber
}

func Run(a *App) error {
	return a.fiber.Listen(":" + a.port)
}
