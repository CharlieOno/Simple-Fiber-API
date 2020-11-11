package models

import (
	"gorm.io/gorm"
)

// Article model
type Article struct {
	gorm.Model
	Title   string
	Author  string
	Content string
}
