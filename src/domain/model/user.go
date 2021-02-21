package model

import (
	"errors"
)

// TODO: Remove gorm meta data from model, because domain must not know technical knowledge.
// However, conversion type is not smart in infrastructure layer. I'm just looking for a good idea.
// User Struct of an user
type User struct {
	ID        int `gorm:"primary_key"`
	LastName  string
	FirstName string
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
