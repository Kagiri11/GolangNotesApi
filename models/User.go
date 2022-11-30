package models

import "time"

type User struct {
	ID        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
