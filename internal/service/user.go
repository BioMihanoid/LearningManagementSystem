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

// TODO: do new fn ChangeUserRole

func (u *User) ChangeUserRole(id int, role string) error {
	return nil
}

// TODO: do new fn UpdateUser

func (u *User) UpdateUser(changeUser models.User) error {

	return nil
}
