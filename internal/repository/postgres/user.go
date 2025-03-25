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
	query := fmt.Sprintf("SELECT user_id, first_name, last_name, email, role_id FROM %s", usersTable)

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)
	for rows.Next() {
		u := models.User{}
		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.RoleID)
		users = append(users, u)
	}

	return users, err
}

func (u *User) GetUserByID(id int) (models.User, error) {
	query := fmt.Sprintf("SELECT first_name, last_name, email, role_id FROM %s WHERE user_id = $1", usersTable)

	rows, err := u.db.Query(query, id)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	for rows.Next() {
		err = rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.RoleID)
	}
	if err != nil {
		return models.User{}, err
	}

	user.ID = id

	return user, err
}

// TODO: do new fn ChangeUserRole

func (u *User) ChangeUserRole(id int, role string) error {
	return nil
}

// TODO: do new fn UpdateUser

func (u *User) UpdateUser(user models.User) error {
	return nil
}
