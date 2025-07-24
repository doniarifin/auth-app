package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := User{}

	result := db.Select("id", "email", "password").Where(&User{Email: email}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	result := db.Select("id", "email").First(&User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return []User{}, nil
}
