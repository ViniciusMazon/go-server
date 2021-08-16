package migrations

import (
	"github.com/viniciusmazon/go-server/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
