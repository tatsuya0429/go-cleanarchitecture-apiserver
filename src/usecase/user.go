package usecase

import (
	"go-rest/domain/model"
	"go-rest/domain/repository"
)

// UserUsecase Interface of an user usecase
type UserUsecase interface {
	Create(lastname, firstname string) (*model.ModelUser, error)
	ReadByID(id int) (*model.ModelUser, error)
	ReadAll() (*model.ModelUsers, error)
	Update(id int, lastname, firstname string) (*model.ModelUser, error)
	Delete(id int) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUsecase Constructor of an user usecase
func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

// Create Usecase to save an user
func (userUsecase *userUsecase) Create(lastname, firstname string) (*model.ModelUser, error) {
	user, err := model.NewUser(lastname, firstname)
	if err != nil {
		return nil, err
	}

	createdUser, err := userUsecase.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// ReadByID Usecase to read an user by id
func (userUsecase *userUsecase) ReadByID(id int) (*model.ModelUser, error) {
	foundUser, err := userUsecase.userRepository.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

// ReadAll Usecase to read users
func (userUsecase *userUsecase) ReadAll() (*model.ModelUsers, error) {
	foundUsers, err := userUsecase.userRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundUsers, nil
}

// Update Usecase to update an user
func (userUsecase *userUsecase) Update(id int, lastname, firstname string) (*model.ModelUser, error) {
	targetUser, err := userUsecase.userRepository.ReadByID(id)
	if err != nil {
		return nil, err
	}

	if err := targetUser.SetUser(lastname, firstname); err != nil {
		return nil, err
	}

	updatedUser, err := userUsecase.userRepository.Update(targetUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// Delete Usecase to delete an user
func (userUsecase *userUsecase) Delete(id int) error {
	targetUser, err := userUsecase.userRepository.ReadByID(id)
	if err != nil {
		return err
	}

	err = userUsecase.userRepository.Delete(targetUser)
	if err != nil {
		return err
	}

	return nil
}
