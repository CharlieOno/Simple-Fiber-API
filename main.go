package main

import (
	"api/config"
	"github.com/gofiber/fiber/v2"
	"api/router"
)

func main() {
	app := fiber.New()

	config.InitLogger(app)
	config.InitDatabase()
	router.SetupRoutes(app)
	router.SetupAuthenticatedRoutes(app)
	err := app.Listen(":4040")
	if err != nil {
		panic(err)
	}
}