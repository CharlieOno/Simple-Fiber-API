package handlers

import (
	"github.com/gofiber/fiber/v2"
	"api/database"
	"api/models"
)

// GetArticle [GET '/api/article/:id'] : get single article by ID from DB
func GetArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article

	if err := database.DBConn.First(&article, id).Error; err != nil {
		return c.Status(500).SendString("Internal server error.")
	}
	c.JSON(article)
	return c.SendStatus(200)
}

// GetArticles [GET '/api/article'] : get articles from DB
func GetArticles(c *fiber.Ctx) error {
	var articles []models.Article

	if err := database.DBConn.Find(&articles).Error; err != nil {
		return c.Status(500).SendString("Internal server error.")
	}
	c.JSON(articles)
	return c.SendStatus(200)
}

// AddArticle [POST '/api/article'] : add an article to DB
func AddArticle(c *fiber.Ctx) error {
	var article models.Article

	if err := c.BodyParser(&article); err != nil {
		return c.Status(503).SendString("Failing during parsing.")
	}
	if err := database.DBConn.Create(&article).Error; err != nil {
		return c.Status(500).SendString("Internal server error.")
	}
	c.JSON(article)
	return c.SendStatus(200)
}

// UpdateArticle [PUT '/api/article/:id'] : update an article from DB
func UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article

	if err := database.DBConn.First(&article, id).Error; err != nil {
		return c.Status(500).SendString("Internal server error.")
	}
	if err := c.BodyParser(&article); err != nil {
		return c.Status(503).SendString("Failing during parsing.")
	}
	database.DBConn.Save(&article)
	c.JSON(article)
	return c.SendStatus(200)
}

// DeleteArticle [DELETE '/api/article/:id'] : delete an article by ID from DB
func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article

	if err := database.DBConn.First(&article, id).Error; err != nil {
		return c.Status(500).SendString("Internal server error.")
	}
	database.DBConn.Delete(&article)
	return c.SendStatus(200)
}


