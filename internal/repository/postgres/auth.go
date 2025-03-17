package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/models"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) CreateUser(user models.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, role) values($1, $2, $3, $4) RETURNING id", usersTable)

	a.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role)

	return a.GetUserID(user.Email)
}

func (a *Auth) GetUserID(email string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE email = $1", usersTable)
	row := a.db.QueryRow(query, email)
	if row.Scan(&id) == nil {
		return id, nil
	}
	return id, nil
}

func (a *Auth) GetUser(email string, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s where  email = $1", usersTable)

	rows, err := a.db.Query(query, email)
	if err != nil {
		return models.User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
