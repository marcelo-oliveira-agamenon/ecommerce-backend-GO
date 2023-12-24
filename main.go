package main

import (
	"log"
	"os"

	"github.com/ecommerce/adapters/primary"
	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/adapters/secondary/token/jwt"
	"github.com/ecommerce/core/services/users"
	_ "github.com/pdrum/swagger-automation/docs"
)

func main()  {
	postgresRepository, err := postgres.NewPostgresRepository()
	if err != nil {
		log.Fatal()
	}

	tokenService := jwt.NewToken(os.Getenv("JWT_KEY"))

	userRepository := postgres.NewUserRepository(postgresRepository)
	userService := users.NewUserService(userRepository)

	srv := primary.NewApp(tokenService, userService, os.Getenv("PORT"))
	primary.Run(srv)
}