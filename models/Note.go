package models

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID          int32     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:100;not null" json:"title"`
	Description string    `gorm:"size:255;not null" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (n *Note) CreateNote(db *gorm.DB) (*Note, error) {
	return nil, nil
}

func (n *Note) GetNote() (*Note, error) {
	return nil, nil
}

func (n *Note) GetNotes() (*[]Note, error) {
	return nil, nil
}

func (n *Note) UpdateNote() (*Note, error) {
	return nil, nil
}

func (n *Note) DeleteNote() error {
	return nil
}
