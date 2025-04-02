package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
)

type Enrollment struct {
	repos *repository.Repository
}

func NewEnrollment(repos *repository.Repository) *Enrollment {
	return &Enrollment{
		repos: repos,
	}
}

func (e *Enrollment) SubscribeOnCourse(userID int, CourseID int) error {
	return e.repos.SubscribeOnCourse(userID, CourseID)
}

func (e *Enrollment) UnsubscribeOnCourse(userID int, CourseID int) error {
	return e.repos.UnsubscribeOnCourse(userID, CourseID)
}

func (e *Enrollment) GetAllUserSubscribedToTheCourse(courseID int) ([]models.Enrollment, error) {
	return e.repos.GetAllUserSubscribedToTheCourse(courseID)
}

func (e *Enrollment) GetAllCoursesCurrentUser(userID int) ([]models.Enrollment, error) {
	return e.repos.GetAllCoursesCurrentUser(userID)
}
