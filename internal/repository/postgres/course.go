package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
)

type Course struct {
	db *sql.DB
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) CreateCourse(course models.Course) error {
	query := fmt.Sprintf("INSERT INTO %s(title, description) VALUES($1, $2)")
	_, err := c.db.Exec(query, course.Title, course.Description)
	return err
}

func (c *Course) GetCourseByID(id int) (models.Course, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE course_id=$1", courseTable)
	row := c.db.QueryRow(query, id)
	course := models.Course{}
	err := row.Scan(&course.ID, &course.Title, &course.Title, &course.CreatedAt)
	if err != nil {
		return course, err
	}
	return course, nil
}

func (c *Course) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY course_id ASC", courseTable)
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var course models.Course
		err = rows.Scan(&course.ID, &course.Title, &course.Description, &course.CreatedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *Course) UpdateTitleCourse(id int, title string) error {
	query := fmt.Sprintf("UPDATE %s set title = $1 where course_id =  $2", usersTable)
	_, err := c.db.Exec(query, title, id)
	return err
}

func (c *Course) UpdateDescriptionCourse(id int, description string) error {
	query := fmt.Sprintf("UPDATE %s SET description = $1 WHERE course_id =  $2", usersTable)
	_, err := c.db.Exec(query, description, id)
	return err
}

func (c *Course) DeleteCourseByID(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", courseTable)
	_, err := c.db.Exec(query, id)
	return err
}
