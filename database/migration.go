package database

import (
	"auth-app/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
