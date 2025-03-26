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
	ChangeUserRole(id int, roleId int) error
	UpdateFirstName(id int, name string) error
	UpdateLastName(id int, name string) error
	UpdateEmail(id int, email string) error
	DeleteUser(id int) error
}

type RoleRepository interface {
	GetLevelAccess(id int) (int, error)
}

type Repository struct {
	Authorization
	UserRepository
	RoleRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:  postgres.NewAuth(db),
		UserRepository: postgres.NewUser(db),
		RoleRepository: postgres.NewRole(db),
	}
}
