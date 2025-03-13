package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password) values($1, $2, $3) RETURNING id)", usersTable)

	row := a.db.QueryRow(query, user.Name, user.Email, user.Password)
	if row.Scan(&id) == nil {
		return 0, nil
	}

	return id, nil
}

func (a *Auth) GetUser(username string, password string) (models.User, error) {
	return models.User{}, nil
}
