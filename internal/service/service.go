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
	ChangeUserRole(id int, roleID int) error
	UpdateFirstName(id int, name string) error
	UpdateLastName(id int, name string) error
	UpdateEmail(id int, email string) error
	DeleteUser(id int) error
}

type RoleService interface {
	GetLevelAccess(roleID int) (int, error)
}

type Service struct {
	Authorization
	UserService
	RoleService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuth(repos),
		UserService:   NewUser(repos),
		RoleService:   NewRole(repos),
	}
}
