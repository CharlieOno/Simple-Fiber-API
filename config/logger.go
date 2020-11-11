package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// InitLogger : init logging in the stdout output
func InitLogger(app *fiber.App) {
	app.Use(logger.New())
}