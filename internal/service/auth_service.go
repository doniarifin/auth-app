package service

import (
	"auth-app/internal/dto"
	"auth-app/internal/model"
	"auth-app/internal/pkg/jwt"
	"auth-app/internal/repository"
	"auth-app/internal/utils"
	"errors"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{r}
}

func (s UserService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, _ := utils.HashPassword(req.Password)

	user := model.User{
		ID:       utils.GenerateUUID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s UserService) Login(req *dto.LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		return "", errors.New("error generate token")
	}

	return token, nil
}
