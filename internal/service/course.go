package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Course struct {
	repos *repository.Repository
}

func NewCourse(repos *repository.Repository) *Course {
	return &Course{
		repos: repos,
	}
}

func (c *Course) CreateCourse(course models.Course) error {
	return c.repos.CreateCourse(course)
}

func (c *Course) UpdateTitleCourse(id int, title string) error {
	return c.repos.UpdateTitleCourse(id, title)
}

func (c *Course) UpdateDescriptionCourse(id int, description string) error {
	return c.repos.UpdateDescriptionCourse(id, description)
}

func (c *Course) GetCourseByID(id int) (models.Course, error) {
	return c.repos.GetCourseByID(id)
}

func (c *Course) GetAllCourses() ([]models.Course, error) {
	return c.repos.GetAllCourses()
}

func (c *Course) DeleteCourseByID(id int) error {
	return c.repos.DeleteCourseByID(id)
}
