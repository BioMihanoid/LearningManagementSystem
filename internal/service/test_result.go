package service

import (
	"errors"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type TestResult struct {
	repos *repository.Repository
}

func NewTestResult(repos *repository.Repository) *TestResult {
	return &TestResult{repos: repos}
}

func (t *TestResult) CreateTestResult(userID int, testID int, score int) error {
	return t.repos.CreateTestResult(userID, testID, score)
}

func (t *TestResult) GetTestResult(userID int, testResultID int) (models.TestResult, error) {
	testResult, err := t.repos.GetTestResult(testResultID)
	if err != nil {
		return models.TestResult{}, err
	}
	if testResult.UserID != userID {
		return models.TestResult{}, errors.New("user is not the right user")
	}
	return testResult, nil
}

func (t *TestResult) UpdateTestResult(testResultID int, score int) error {
	return t.repos.UpdateTestResult(testResultID, score)
}

func (t *TestResult) DeleteTestResult(testResultID int) error {
	return t.repos.DeleteTestResult(testResultID)
}
