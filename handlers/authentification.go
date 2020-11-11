package handlers

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login : retrieve JWT for authorization
func Login(c *fiber.Ctx) error {
	var body credentials

	if err := c.BodyParser(&body); err != nil {
		return c.Status(503).SendString("Failing during parsing.")
	}

	// Retrieve user based on credentials
	if body.Email != "john" || body.Password != "doe" {
		return c.Status(401).SendString("Unauthorized.")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = "john"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(500).SendString("Internal server error.")
	}

	return c.JSON(fiber.Map{
		"token": t,
		"user": struct {
			ID int `json:"id"`
			Email string `json:"email"`
		}{
			ID: 1,
			Email: "john",
		},
	})
}

// Whoami : retrieve user informations based on token
func Whoami(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.JSON(fiber.Map{
		"user" : claims["user"].(string),
		"admin" : claims["admin"].(bool),
	})
}