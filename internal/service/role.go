package service

import "github.com/BioMihanoid/LearningManagementSystem/internal/repository"

type Role struct {
	repos *repository.Repository
}

func NewRole(repos *repository.Repository) *Role {
	return &Role{
		repos: repos,
	}
}

func (r *Role) GetLevelAccess(roleID int) (int, error) {
	return r.repos.GetLevelAccess(roleID)
}
