package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/ecommerce/adapters/primary"
	"github.com/ecommerce/adapters/secondary/email/gomail"
	"github.com/ecommerce/adapters/secondary/postgres"
	storage "github.com/ecommerce/adapters/secondary/storage/aws"
	"github.com/ecommerce/adapters/secondary/token/jwt"
	categories "github.com/ecommerce/core/services/category"
	favorites "github.com/ecommerce/core/services/favorite"
	productImages "github.com/ecommerce/core/services/productImage"
	"github.com/ecommerce/core/services/products"
	"github.com/ecommerce/core/services/users"
	"github.com/joho/godotenv"
	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	godotenv.Load(".env")
	postgresRepository, err := postgres.NewPostgresRepository()
	if err != nil {
		log.Fatal() //todo: maybe change this?
	}

	jtwKey := os.Getenv("JWT_KEY")
	port := os.Getenv("PORT")
	config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_SECRET_ID"), os.Getenv("AWS_SECRET_KEY"), ""),
	}
	storageService := storage.NewAWS(*config)
	tokenService := jwt.NewToken(jtwKey)
	emailService := gomail.NewEmailService()

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

	srv := primary.NewApp(tokenService, storageService, userService, productService, categoryService, productImageService, favoriteService, emailService, port)
	primary.Run(srv)
}
