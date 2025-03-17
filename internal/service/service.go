package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(name string, password string) (models.User, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuth(repos),
	}
}
