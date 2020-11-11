package config

import (
	"gorm.io/gorm"
	"api/models"
  	"gorm.io/driver/sqlite"
)

// DB : Gorm database instance
var DB *gorm.DB

// InitDatabase : initiate SQLite database with auto-migrations
func InitDatabase() {
	var err error

	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}
	DB.AutoMigrate(&models.Article{})
}