package models

import "time"

type User struct {
	ID           int
	Name         string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
