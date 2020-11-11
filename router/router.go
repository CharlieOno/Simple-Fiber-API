package router

import (
	"github.com/gofiber/fiber/v2"
	"api/handlers"
	"api/middlewares"
)

// SetupAuthenticatedRoutes : setup authenticated routes
func SetupAuthenticatedRoutes(app *fiber.App) {
	middlewares.JWT(app);
	api := app.Group("/api", middlewares.SetSecurityHeaders, middlewares.CORS)

	// Authentification
	api.Get("/whoami", handlers.Whoami)

	// Articles
	api.Get("/article/:id", handlers.GetArticle)
	api.Get("/article", handlers.GetArticles)
	api.Put("/article/:id", handlers.UpdateArticle)
	api.Delete("/article/:id", handlers.DeleteArticle)
	api.Post("/article", handlers.AddArticle)
}

// SetupRoutes : setup unauthenticated routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/", middlewares.SetSecurityHeaders, middlewares.CORS)

	// Informations
	api.Get("/", handlers.GetInformations)
	api.Post("/login", handlers.Login)
}