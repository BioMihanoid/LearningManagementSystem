package postgres

import (
	"database/sql"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) CreateUser(user models.User) error {
	return nil
}

func (a *Auth) GetUser(username string, password string) (models.User, error) {
	return models.User{}, nil
}
