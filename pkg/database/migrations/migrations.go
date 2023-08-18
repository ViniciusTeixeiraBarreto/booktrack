package migrations

import (
	"booktrack/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		models.Author{},
		models.Book{},
	)
}
