package model

import (
	"errors"
)

// User Struct of an user
type User struct {
	ID        int    `gorm:"primary_key" json:"id"`
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
}

type Users []User

// NewUser Constructor of an user
func NewUser(lastname, firstname string) (*User, error) {
	if lastname == "" {
		return nil, errors.New("性は必須です。")
	}

	user := &User{
		LastName:  lastname,
		FirstName: firstname,
	}

	return user, nil
}

// SetUser Setter of an User
func (user *User) SetUser(lastname, firstname string) error {
	if lastname == "" {
		return errors.New("性は必須です。")
	}

	user.LastName = lastname
	user.FirstName = firstname

	return nil
}
