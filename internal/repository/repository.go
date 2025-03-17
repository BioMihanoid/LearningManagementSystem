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

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuth(db),
	}
}
