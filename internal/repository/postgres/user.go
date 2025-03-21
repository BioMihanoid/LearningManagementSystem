package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) GetAllUsers() ([]models.User, error) {
	query := fmt.Sprintf("SELECT id, username, email, role FROM %s", usersTable)

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)
	for rows.Next() {
		u := models.User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role)
		users = append(users, u)
	}

	return users, err
}

func (u *User) GetUserByID(id int) (models.User, error) {
	query := fmt.Sprintf("SELECT username, email, role FROM %s WHERE id = $1", usersTable)

	rows, err := u.db.Query(query, id)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	for rows.Next() {
		err = rows.Scan(&user.Name, &user.Email, &user.Role)
	}
	if err != nil {
		return models.User{}, err
	}

	user.ID = uint(id)

	return user, err
}

func (u *User) ChangeUserRole(id int, role string) error {
	query := fmt.Sprintf("UPDATE %s SET role = $1 WHERE id = $2", usersTable)
	_, err := u.db.Exec(query, role, id)
	return err
}

func (u *User) UpdateUser(user models.User) error {
	query := fmt.Sprintf("UPDATE %s set username = $1 WHERE id = $2", usersTable)
	_, err := u.db.Exec(query, user.Name, int(user.ID))
	query = fmt.Sprintf("UPDATE %s set email = $1 WHERE id = $2", usersTable)
	_, err = u.db.Exec(query, user.Email, int(user.ID))
	return err
}
