package models

type Lesson struct {
	ID       int
	CourseID int
	Title    string
	Content  string
	TestId   int
}
