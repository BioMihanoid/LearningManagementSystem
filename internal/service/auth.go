package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Auth struct {
	repo *repository.Authorization
}

func NewAuth(repo *repository.Authorization) *Auth {
	return &Auth{}
}

func (a *Auth) CreateUser(user models.User) error {
	return nil
}

func (a *Auth) GetUser(name string, passwordHash string) (models.User, error) {
	return models.User{}, nil
}
