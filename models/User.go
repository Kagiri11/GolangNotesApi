package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Model(&User{}).Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
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
