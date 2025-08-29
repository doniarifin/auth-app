package repository

import (
	"auth-app/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByEmail(string) (*model.User, error)
	Create(model.User) error
}

type userRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryGorm{db}
}

func (u userRepositoryGorm) FindAll() ([]model.User, error) {
	result := u.db.Find([]model.User{}, model.User{})
	if result.Error != nil {
		return nil, result.Error
	}
	return []model.User{}, nil
}

func (u userRepositoryGorm) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where(&model.User{Email: email}).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (u userRepositoryGorm) Create(m model.User) error {
	return u.db.Create(m).Error
}
