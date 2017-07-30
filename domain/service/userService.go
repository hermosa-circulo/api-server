package service

import (
	"errors"
	"iga-controller/domain/models"
	"iga-controller/domain/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repository repository.userRepository.GetUserRepository()
		},
	}
}

func (s *UserService) GetUserService() []*models.User {
	return s.repository.GetUsers();
}

func (s *UserService) GetUserByIdService(id int) (*models.User, error) {
	return s.repository.GetUserById(id)
}

func (s *UserService) SaveUserService(user *models.User) error {
	return s.repository.SaveUser(user)
}

var userService *UserService

func GetUserService() (s UserService) {
	if userService == nil {
		userService = NewUserService()
	}
	return userService
}
