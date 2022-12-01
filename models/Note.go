package models

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID          int32     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:100;not null" json:"title"`
	Description string    `gorm:"size:255;not null" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (n *Note) CreateNote(db *gorm.DB) (*Note, error) {
	err := db.Debug().Model(&Note{}).Create(&n).Error
	if err != nil {
		return &Note{}, err
	}
	return n, nil
}

func (n *Note) GetNote(db *gorm.DB, noteId int) (*Note, error) {
	err := db.Debug().Find(&Note{}).Where("id = ?").Take(&n).Error
	if err != nil {
		return &Note{}, err
	}
	return n, nil
}

func (n *Note) GetNotes(db *gorm.DB) (*[]Note, error) {
	var notes []Note
	err := db.Debug().Model(&Note{}).Limit(100).Find(&notes).Error
	if err != nil {
		return &[]Note{}, err
	}
	return &notes, nil
}

func (n *Note) UpdateNote(db *gorm.DB) (*Note, error) {
	return nil, nil
}

func (n *Note) DeleteNote(db *gorm.DB, noteId int) error {
	err := db.Debug().Where("id = ?", noteId).Delete(&n).Error
	if err != nil {
		return err
	}
	return nil
}
