package main

import (
	"github.com/gofiber/fiber/v2"
	"api/handlers"
	"api/middlewares"
	"api/models"
	"api/database"
	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api", middlewares.SetSecurityHeaders, middlewares.CORS)

	// Informations
	api.Get("/", handlers.GetInformations)

	// Articles
	api.Get("/article/:id", handlers.GetArticle)
	api.Get("/article", handlers.GetArticles)
	api.Put("/article/:id", handlers.UpdateArticle)
	api.Delete("/article/:id", handlers.DeleteArticle)
	api.Post("/article", handlers.AddArticle)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}
	database.DBConn.AutoMigrate(&models.Article{})
}

func main() {
  	app := fiber.New()

	initDatabase()
	setupRoutes(app)
  	app.Listen(":4040")
}