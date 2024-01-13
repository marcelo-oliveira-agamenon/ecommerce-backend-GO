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
	"github.com/ecommerce/core/services/users"
	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	postgresRepository, err := postgres.NewPostgresRepository()
	if err != nil {
		log.Fatal()
	}

	config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_SECRET_ID"), os.Getenv("AWS_SECRET_KEY"), ""),
	}
	storageService := storage.NewAWS(*config)

	tokenService := jwt.NewToken(os.Getenv("JWT_KEY"))

	emailService := gomail.NewEmailService()

	userRepository := postgres.NewUserRepository(postgresRepository)
	userService := users.NewUserService(userRepository)

	srv := primary.NewApp(tokenService, storageService, userService, emailService, os.Getenv("PORT"))
	primary.Run(srv)
}
