package models

import "time"

type User struct {
	ID        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
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
