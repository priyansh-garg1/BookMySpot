package db

import (
	"github.com/priyansh-garg1/ticket-verse/models"
	"gorm.io/gorm"
)

func DBMigrater(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{
		
	})
}
