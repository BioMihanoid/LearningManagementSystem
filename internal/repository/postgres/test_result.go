package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
)

type TestResult struct {
	db *sql.DB
}

func NewTestResult(db *sql.DB) *TestResult {
	return &TestResult{db: db}
}

func (t *TestResult) CreateTestResult(userID int, testID int, score int) error {
	query := fmt.Sprintf("INSERT INTO %s(user_id,  test_id, score) VALUES($1, $2, $3)", testResultsTable)
	_, err := t.db.Exec(query, userID, testID, score)
	return err
}

func (t *TestResult) GetTestResult(testResultID int) (models.TestResult, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE test_result_id=$1", testResultsTable)
	var testResult models.TestResult
	err := t.db.QueryRow(query, testResultID).Scan(
		&testResult.UserID,
		&testResult.TestID,
		&testResult.Score)
	return testResult, err
}

func (t *TestResult) UpdateTestResult(testResultID int, score int) error {
	query := fmt.Sprintf("UPDATE %s SET score=$1 WHERE test_result_id=$2", testResultsTable)
	_, err := t.db.Exec(query, score, testResultID)
	return err
}

func (t *TestResult) DeleteTestResult(testResultID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE test_result_id=$1", testResultsTable)
	_, err := t.db.Exec(query, testResultID)
	return err
}
