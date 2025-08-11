package models

import (
	"auth-app/utils"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey" json:"id"`
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

func UpdateUser(db *gorm.DB, user *User, id string) error {
	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(db *gorm.DB, id string) error {
	result := db.Where(&User{ID: id}).Delete(&User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
