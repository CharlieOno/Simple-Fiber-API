package middlewares

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)

// CORS : handling Cross-Origin Resource Sharing with custom configuration
func CORS(c *fiber.Ctx) error {
	cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders:  "Origin, Content-Type, Accept",
	})

	return c.Next()
}