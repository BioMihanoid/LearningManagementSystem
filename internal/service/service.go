package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(name string, password string) (models.User, error)
}

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	ChangeUserRole(id int, role string) error
}

type Service struct {
	Authorization
	UserService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuth(repos),
		UserService:   NewUser(repos),
	}
}
