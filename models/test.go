package models

import "time"

type Test struct {
	ID        int       `json:"id"`
	CourseID  int       `json:"course_id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
}
