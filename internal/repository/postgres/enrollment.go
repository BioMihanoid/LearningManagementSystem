package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Enrollment struct {
	db *sql.DB
}

func NewEnrollment(db *sql.DB) *Enrollment {
	return &Enrollment{db: db}
}

func (e *Enrollment) SubscribeOnCourse(userID int, courseID int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, course_id) VALUES ($1, $2)", enrollmentsTable)
	_, err := e.db.Exec(query, userID, courseID)
	return err
}

func (e *Enrollment) UnsubscribeOnCourse(userID int, courseID int) error {
	query := fmt.Sprintf("DELETE FROM %S WHERE user_id=$1 AND course_id=$2", enrollmentsTable)
	_, err := e.db.Exec(query, userID, courseID)
	return err
}

func (e *Enrollment) GetAllUserSubscribedToTheCourse(courseID int) ([]models.Enrollment, error) {
	query := fmt.Sprintf("SELECT * FROM %s where course_id=$1", enrollmentsTable)
	var enrollments []models.Enrollment
	rows, err := e.db.Query(query, courseID)
	if err != nil {
		return enrollments, err
	}
	defer rows.Close()
	for rows.Next() {
		var enrollment models.Enrollment
		err = rows.Scan(&enrollment.ID, enrollment.CourseID, enrollment.UserID)
		if err != nil {
			return enrollments, err
		}
	}
	return enrollments, nil
}

func (e *Enrollment) GetAllCoursesCurrentUser(userID int) ([]models.Enrollment, error) {
	query := fmt.Sprintf("SELECT * FROM %s where user_id=$1", enrollmentsTable)
	var enrollments []models.Enrollment
	rows, err := e.db.Query(query, userID)
	if err != nil {
		return enrollments, err
	}
	defer rows.Close()
	for rows.Next() {
		var enrollment models.Enrollment
		err = rows.Scan(&enrollment.ID, enrollment.CourseID, enrollment.UserID)
		if err != nil {
			return enrollments, err
		}
	}
	return enrollments, nil
}
