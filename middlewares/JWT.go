package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// JWT : handling JWT token to authenticate user
func JWT(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("supersecret"),
	}))
}