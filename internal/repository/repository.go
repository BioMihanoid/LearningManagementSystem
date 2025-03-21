package repository

import (
	"database/sql"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository/postgres"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email string, passwordHash string) (models.User, error)
	GetUserID(email string) (int, error)
}

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	ChangeUserRole(id int, role string) error
	UpdateUser(user models.User) error
}

type Repository struct {
	Authorization
	UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:  postgres.NewAuth(db),
		UserRepository: postgres.NewUser(db),
	}
}
