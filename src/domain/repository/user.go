package repository

import (
	"go-rest/domain/model"
)

// UserRepository Interface of an user repository
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	ReadByID(id int) (*model.User, error)
	ReadAll() (*model.Users, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) error
}
