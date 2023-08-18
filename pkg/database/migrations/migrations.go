package migrations

import (
	"web-api/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		models.Author{},
		models.Book{},
	)
}
