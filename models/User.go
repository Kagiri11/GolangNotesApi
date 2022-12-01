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

func (u *User) GetUser(db *gorm.DB, id int) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func (u *User) UpdateUser() (*User, error) {
	return nil, nil
}

func (u *User) DeleteUser(db *gorm.DB, id int) error {
	err := db.Debug().Where("id = ?", id).Delete(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) LoginUser() error {
	return nil
}
