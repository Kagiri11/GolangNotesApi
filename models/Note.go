package models

import "time"

type Note struct {
	ID          int32
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
