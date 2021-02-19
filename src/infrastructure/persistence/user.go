package persistence

import (
	"go-rest/domain/model"
	"go-rest/domain/repository"

	"github.com/jinzhu/gorm"
)

// UserRepository struct of an user repository
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository Constructor of an user repository
func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create Create an user
func (userRepository *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := userRepository.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// ReadByID Read an user by ID
func (userRepository *UserRepository) ReadByID(id int) (*model.User, error) {
	user := &model.User{ID: id}

	if err := userRepository.Conn.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// ReadAll Read users
func (userRepository *UserRepository) ReadAll() (*model.Users, error) {
	users := &model.Users{}

	// gorm.Find from v2 doesn't return ErrRecordNotFound
	userRepository.Conn.Find(&users)

	return users, nil
}

// Update Update an user
func (userRepository *UserRepository) Update(user *model.User) (*model.User, error) {
	if err := userRepository.Conn.Model(&user).Update(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Delete Delete an user
func (userRepository *UserRepository) Delete(user *model.User) error {
	if err := userRepository.Conn.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
