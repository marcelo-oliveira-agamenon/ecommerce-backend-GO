package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/ecommerce/adapters/primary"
	"github.com/ecommerce/adapters/secondary/email/gomail"
	kafka_ins "github.com/ecommerce/adapters/secondary/kafka"
	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/adapters/secondary/redis"
	storage "github.com/ecommerce/adapters/secondary/storage/aws"
	cronjob "github.com/ecommerce/adapters/secondary/tasks/cron"
	"github.com/ecommerce/adapters/secondary/token/jwt"
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
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	postgresRepository, errP := postgres.NewPostgresRepository()
	if errP != nil {
		log.Fatal(errP)
	}
	redisRepository, errR := redis.NewRedisRepository()
	if errR != nil {
		log.Fatal(errR)
	}
	kafkaRepository, errK := kafka_ins.NewKafkaRepository()
	if errK != nil {
		log.Fatal(errK)
	}
	cronjob.NewCronTasks(postgresRepository)

	jtwKey := os.Getenv("JWT_KEY")
	port := os.Getenv("PORT")
	config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_SECRET_ID"), os.Getenv("AWS_SECRET_KEY"), ""),
	}
	storageService := storage.NewAWS(*config)
	tokenService := jwt.NewToken(jtwKey)
	emailService := gomail.NewEmailService()
	redisService := redis.NewRedisSessionRepository(redisRepository)
	kafkaService := kafka_ins.NewKafkaSessionRepository(kafkaRepository)

	userRepository := postgres.NewUserRepository(postgresRepository)
	userService := users.NewUserService(userRepository)
	productRepository := postgres.NewProductRepository(postgresRepository)
	productService := products.NewProductService(productRepository)
	categoryRepository := postgres.NewCategoryRepository(postgresRepository)
	categoryService := categories.NewCategoryService(categoryRepository)
	productImageRepository := postgres.NewProductImageRepository(postgresRepository)
	productImageService := productImages.NewProductImageService(productImageRepository)
	favoriteRepository := postgres.NewFavoriteRepository(postgresRepository)
	favoriteService := favorites.NewFavoriteService(favoriteRepository)
	couponRepository := postgres.NewCouponRepository(postgresRepository)
	couponService := coupons.NewCouponService(couponRepository)
	orderRepository := postgres.NewOrderRepository(postgresRepository)
	orderService := orders.NewOrderService(orderRepository)
	orderDetailsRepository := postgres.NewOrderDetailsRepository(postgresRepository)
	orderDetailsService := ordersdetails.NewOrdersDetailsService(orderDetailsRepository)
	paymentRepository := postgres.NewPaymentRepository(postgresRepository)
	paymentService := payments.NewPaymentService(paymentRepository)
	logRepository := postgres.NewLogRepository(postgresRepository)
	logService := logs.NewLogService(logRepository)
	miscRepository := postgres.NewMiscRepository(postgresRepository)
	miscService := misc.NewMiscService(miscRepository)

	srv := primary.NewApp(
		tokenService, storageService, userService, productService, categoryService,
		productImageService, favoriteService, couponService, orderService,
		orderDetailsService, paymentService, logService, miscService, emailService,
		redisService, kafkaService, port)
	primary.Run(srv)
}
