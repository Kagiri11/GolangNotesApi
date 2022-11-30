package models

import "time"

type User struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) CreateUser() (*User, error) {
	return nil, nil
}

func (u *User) GetUser() (*User, error) {
	return nil, nil
}

func (u *User) GetUsers() (*[]User, error) {
	return nil, nil
}

func (u *User) UpdateUser() (*User, error) {
	return nil, nil
}

func (u *User) DeleteUser() error {
	return nil
}

func (u *User) LoginUser() error {
	return nil
}
