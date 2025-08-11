package logics

import (
	"auth-app/models"
	"errors"

	"gorm.io/gorm"
)

func FindUserByID(db *gorm.DB, id string) (*models.User, error) {
	user := models.User{}

	result := db.Select("id", "email", "password").Where(&models.User{ID: id}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	user := models.User{}

	result := db.Select("id", "email", "password").Where(&models.User{Email: email}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	result := db.Select("id", "email").First(&models.User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return []models.User{}, nil
}
