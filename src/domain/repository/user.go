package repository

import (
	"go-rest/domain/model"
)

// UserRepository Interface of an user repository
type UserRepository interface {
	Create(user *model.ModelUser) (*model.ModelUser, error)
	ReadByID(id int) (*model.ModelUser, error)
	ReadAll() (*model.ModelUsers, error)
	Update(user *model.ModelUser) (*model.ModelUser, error)
	Delete(user *model.ModelUser) error
}
