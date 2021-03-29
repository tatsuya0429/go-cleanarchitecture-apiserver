package persistence

import (
	"go-rest/domain/model"
	"go-rest/domain/repository"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

// User Struct of an user
type User struct {
	ID int `gorm:"primary_key"`
	model.ModelUser
}

// Users Struct of users
type Users []User

// UserRepository struct of an user repository
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository Constructor of an user repository
func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create Create an user
func (userRepository *UserRepository) Create(modelUser *model.ModelUser) (*model.ModelUser, error) {
	user := User{}
	copier.Copy(&user, &modelUser)
	if err := userRepository.Conn.Create(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.ModelUser)
	copier.Copy(&userModel, &user)

	return userModel, nil
}

// ReadByID Read an user by ID
func (userRepository *UserRepository) ReadByID(id int) (*model.ModelUser, error) {
	user := User{ID: id}
	if err := userRepository.Conn.First(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.ModelUser)
	copier.Copy(&userModel, &user)

	return userModel, nil
}

// ReadAll Read users
func (userRepository *UserRepository) ReadAll() (*model.ModelUsers, error) {
	users := Users{}
	// gorm.Find from v2 doesn't return ErrRecordNotFound
	userRepository.Conn.Find(&users)
	userModels := new(model.ModelUsers)
	copier.Copy(&userModels, &users)

	return userModels, nil
}

// Update Update an user
func (userRepository *UserRepository) Update(modelUser *model.ModelUser) (*model.ModelUser, error) {
	user := User{}
	copier.Copy(&user, &modelUser)
	if err := userRepository.Conn.Model(&user).Update(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.ModelUser)
	copier.Copy(&userModel, &user)

	return userModel, nil
}

// Delete Delete an user
func (userRepository *UserRepository) Delete(modelUser *model.ModelUser) error {
	user := User{}
	copier.Copy(&user, &modelUser)
	if err := userRepository.Conn.Delete(&user).Error; err != nil {
		return err
	}
	userModel := new(model.ModelUser)
	copier.Copy(&userModel, &user)

	return nil
}
