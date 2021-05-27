package main

import (
	"log"

	m "github.com/ecommerce/db"
	r "github.com/ecommerce/routes"
	u "github.com/ecommerce/utility"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"

	_ "github.com/pdrum/swagger-automation/docs"
)

func main()  {
	m.CreateConnection()
	app := fiber.New()
	app.Use(cors.New())
	r.Routes(app)
	port := u.GetDotEnv("PORT")
	u.CallerAllJobs()

	log.Fatal(app.Listen(port))
}