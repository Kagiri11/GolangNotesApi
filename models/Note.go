package models

import "time"

type Note struct {
	ID          int32
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (n *Note) CreateNote() (*Note, error) {
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
