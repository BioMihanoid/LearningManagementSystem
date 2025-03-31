package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Test struct {
	repos *repository.Repository
}

func NewTest(repos *repository.Repository) *Test {
	return &Test{
		repos: repos,
	}
}

func (t *Test) CreateTest(test models.Test) error {
	return t.repos.CreateTest(test.CourseID, test.Question, test.Answer)
}

func (t *Test) GetTestByID(testID int) (models.Test, error) {
	return t.repos.GetTestByID(testID)
}

func (t *Test) GetAllTestsCourse(courseID int) ([]models.Test, error) {
	return t.repos.GetAllTestsCourse(courseID)
}

func (t *Test) GetAllTests() ([]models.Test, error) {
	return t.repos.GetAllTests()
}

func (t *Test) UpdateQuestionTest(testID int, question string) error {
	return t.repos.UpdateQuestionTest(testID, question)
}

func (t *Test) UpdateAnswerTest(testID int, question string) error {
	return t.repos.UpdateAnswerTest(testID, question)
}

func (t *Test) DeleteTest(testID int) error {
	return t.repos.DeleteTest(testID)
}

func (t *Test) SubmitTest(testAnswer string, userAnswer string) int {
	if testAnswer == userAnswer {
		return 1
	} else {
		return 0
	}
}
