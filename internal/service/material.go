package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
)

type Material struct {
	repos *repository.Repository
}

func NewMaterial(repos *repository.Repository) *Material {
	return &Material{
		repos: repos,
	}
}

func (m *Material) CreateMaterial(material models.Material) error {
	return m.repos.CreateMaterial(material)
}

func (m *Material) GetMaterialByID(materialID int) (models.Material, error) {
	return m.repos.GetMaterialByID(materialID)
}

func (m *Material) GetCourseMaterial(courseID int) ([]models.Material, error) {
	return m.repos.GetCourseMaterial(courseID)
}

func (m *Material) UpdateTitleMaterial(materialID int, title string) error {
	return m.repos.UpdateTitleMaterial(materialID, title)
}

func (m *Material) UpdateContentMaterial(materialID int, content string) error {
	return m.repos.UpdateContentMaterial(materialID, content)
}

func (m *Material) DeleteMaterial(materialID int) error {
	return m.repos.DeleteMaterial(materialID)
}
