package model

import (
	"errors"
)

// ModelUser Struct of an user
type ModelUser struct {
	ID        int `gorm:"primary_key"`
	LastName  string
	FirstName string
}

type ModelUsers []ModelUser

// NewUser Constructor of an user
func NewUser(lastname, firstname string) (*ModelUser, error) {
	if lastname == "" {
		return nil, errors.New("性は必須です。")
	}

	user := &ModelUser{
		LastName:  lastname,
		FirstName: firstname,
	}

	return user, nil
}

// SetUser Setter of an User
func (user *ModelUser) SetUser(lastname, firstname string) error {
	if lastname == "" {
		return errors.New("性は必須です。")
	}

	user.LastName = lastname
	user.FirstName = firstname

	return nil
}
