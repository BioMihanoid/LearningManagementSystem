package models

import "time"

type Material struct {
	ID        int       `json:"id"`
	CourseID  int       `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
