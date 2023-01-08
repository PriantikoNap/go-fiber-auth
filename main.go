package main

import (
	"github.com/PriantikoNap/go-fiber-auth.git/database"
	"github.com/PriantikoNap/go-fiber-auth.git/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")
}
