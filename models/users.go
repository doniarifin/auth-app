package models

import (
	"auth-app/utils"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func CreateUser(db *gorm.DB, user *User) error {
	user.ID = utils.NewUuid()
	result := db.Create(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func UpdateUser(db *gorm.DB, user *User) error {
	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := User{}

	result := db.Select("id", "email", "password").Where(&User{Email: email}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	result := db.Select("id", "email").First(&User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return []User{}, nil
}
