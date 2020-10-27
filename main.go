package main

import (
	"log"

	m "github.com/ecommerce/db"
	r "github.com/ecommerce/routes"
	u "github.com/ecommerce/utility"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func main()  {
	m.CreateConnection()
	app := fiber.New()
	app.Use(cors.New())
	r.Routes(app)
	port := u.GetDotEnv("PORT")

	log.Fatal(app.Listen(port))
}