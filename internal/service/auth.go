package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Auth struct {
	repos *repository.Repository
}

func NewAuth(repos *repository.Repository) *Auth {
	return &Auth{
		repos: repos,
	}
}

func (a *Auth) CreateUser(user models.User) error {
	return nil
}

func (a *Auth) GetUser(name string, passwordHash string) (models.User, error) {
	return models.User{}, nil
}
