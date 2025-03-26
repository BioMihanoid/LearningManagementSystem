package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type User struct {
	repos *repository.Repository
}

func NewUser(repos *repository.Repository) *User {
	return &User{
		repos: repos,
	}
}

func (u *User) GetAllUsers() ([]models.User, error) {
	return u.repos.GetAllUsers()
}

func (u *User) GetUserById(id int) (models.User, error) {
	return u.repos.GetUserByID(id)
}

func (u *User) ChangeUserRole(id int, roleID int) error {
	// TODO: add check role exist
	return u.repos.ChangeUserRole(id, roleID)
}

func (u *User) UpdateFirstName(id int, name string) error {
	return u.repos.UpdateFirstName(id, name)
}

func (u *User) UpdateLastName(id int, name string) error {
	return u.repos.UpdateLastName(id, name)
}

func (u *User) UpdateEmail(id int, email string) error {
	return u.repos.UpdateEmail(id, email)
}

func (u *User) DeleteUser(id int) error {
	return u.repos.DeleteUser(id)
}
