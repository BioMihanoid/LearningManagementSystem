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

func (u *User) ChangeUserRole(id int, roleID int) error {
	query := fmt.Sprintf("UPDATE %s SET role_id = $1 WHERE user_id = $2", usersTable)
	_, err := u.db.Exec(query, roleID, id)
	return err
}

func (u *User) UpdateFirstName(id int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET first_name = $1 WHERE user_id = $2", usersTable)
	_, err := u.db.Exec(query, name, id)
	return err
}

func (u *User) UpdateLastName(id int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET last_name = $1 WHERE user_id = $2", usersTable)
	_, err := u.db.Exec(query, name, id)
	return err
}

func (u *User) UpdateEmail(id int, email string) error {
	var idCheck int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE email = $1", usersTable)
	row := u.db.QueryRow(query, email)
	if row.Scan(&idCheck) != nil {
		return fmt.Errorf("%s already exists", email)
	}
	query = fmt.Sprintf("UPDATE %s SET email = $1 WHERE user_id = $2", usersTable)
	_, err := u.db.Exec(query, email, id)
	return err
}

func (u *User) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersTable)
	_, err := u.db.Exec(query, id)
	return err
}
