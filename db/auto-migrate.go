package db

import (
	"github.com/mwombeki6/e_water-backend/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}