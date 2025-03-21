package service

import (
	"fmt"
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

func (u *User) ChangeUserRole(id int, role string) error {
	if role != "admin" && role != "teacher" && role != "student" {
		return fmt.Errorf("role %s not supported", role)
	}

	return u.repos.ChangeUserRole(id, role)
}

func (u *User) UpdateUser(changeUser models.User) error {
	user, err := u.GetUserById(int(changeUser.ID))
	if err != nil {
		return err
	}
	_ = user

	if changeUser.Email != "" {
		user.Email = changeUser.Email
	}
	if changeUser.Name != "" {
		user.Name = changeUser.Name
	}

	err = u.repos.UpdateUser(user)

	return err
}
