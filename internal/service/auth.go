package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repos *repository.Repository
}

func NewAuth(repos *repository.Repository) *Auth {
	return &Auth{
		repos: repos,
	}
}

func (a *Auth) CreateUser(user models.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(hashedPassword)

	id, err := a.repos.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *Auth) GetUser(name string, passwordHash string) (models.User, error) {
	return models.User{}, nil
}
