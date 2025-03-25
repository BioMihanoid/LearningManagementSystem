package models

import "time"

type TestResult struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	TestID      int       `json:"test_id"`
	Score       int       `json:"score"`
	CompletedAt time.Time `json:"completed_at"`
}
