package models

import "time"

type Enrollment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CourseID  int       `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
}
