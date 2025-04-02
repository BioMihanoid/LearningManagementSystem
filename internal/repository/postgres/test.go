package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
)

type Test struct {
	db *sql.DB
}

func NewTest(db *sql.DB) *Test {
	return &Test{db: db}
}

func (t *Test) CreateTest(courseID int, question string, answer string) error {
	query := fmt.Sprintf("INSERT INTO %s (course_id, question, answer) VALUES($1, $2, $3)", testsTable)
	_, err := t.db.Exec(query, courseID, question, answer)
	return err
}

func (t *Test) GetTestByID(testID int) (models.Test, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", testsTable)
	row := t.db.QueryRow(query, testID)
	var test models.Test
	err := row.Scan(&test.ID, &test.Question, &test.Answer)
	return test, err
}

func (t *Test) GetAllTestsCourse(courseID int) ([]models.Test, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE course_id=$1", testsTable)
	rows, err := t.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tests []models.Test
	for rows.Next() {
		var test models.Test
		err = rows.Scan(&test.ID, &test.Question, &test.Answer)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}
	return tests, nil
}

func (t *Test) GetAllTests() ([]models.Test, error) {
	query := fmt.Sprintf("SELECT * FROM %s", testsTable)
	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tests []models.Test
	for rows.Next() {
		var test models.Test
		err = rows.Scan(&test.ID, &test.Question, &test.Answer)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}
	return tests, nil
}

func (t *Test) UpdateQuestionTest(testID int, question string) error {
	query := fmt.Sprintf("UPDATE %s SET question=$1 WHERE id=$2", testsTable)
	_, err := t.db.Exec(query, question, testID)
	return err
}

func (t *Test) UpdateAnswerTest(testID int, answer string) error {
	query := fmt.Sprintf("UPDATE %s SET answer=$1 WHERE id=$2", testsTable)
	_, err := t.db.Exec(query, answer, testID)
	return err
}

func (t *Test) DeleteTest(testID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", testsTable)
	_, err := t.db.Exec(query, testID)
	return err
}
