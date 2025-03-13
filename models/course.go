package models

import "time"

type Course struct {
	ID          int
	Title       string
	Description string
	CreatedBy   string
	CreatedUp   time.Duration
}
