package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
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

func (a *Auth) GetUser(name string, password string) (models.User, error) {
	user, err := a.repos.GetUser(name, password)
	if err != nil {
		return user, err
	}

	return user, nil
}
