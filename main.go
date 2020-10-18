package main

import (
	"log"

	m "github.com/ecommerce/db"
	r "github.com/ecommerce/routes"
	"github.com/gofiber/fiber"
)

func main()  {
	m.CreateConnection()
	app := fiber.New()
	r.Routes(app)

	log.Fatal(app.Listen(4000))
}