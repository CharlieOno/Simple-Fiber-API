package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetInformations : serve basic informations about the API
func GetInformations(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}