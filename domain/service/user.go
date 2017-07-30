package service

import (
	"errors"
	"iga-controller/domain/models"
	"iga-controller/domain/repository"
)

type UserService interface {
	GetUsers() []*models.User
	GetUserById(id string) (*models.User, error)
	SaveUser(user *models.User) error
}

type DBUserRepository struct {
	users []*models.User
}

func New() *DBUserRepository {
	return &DBUserRepository{
		users: []*models.User{
			&models.User{Name: "Hector"},
			&models.User{Name: "Carlos"},
			&models.User{Name: "Javi"},
			&models.User{Name: "Dani"},
		},
	}
}

func (r *DBUserRepository) GetUsers() []*models.User {
	return r.users
}

func (r *DBUserRepository) GetUserById(id string) (*models.User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *DBUserRepository) SaveUser(user *models.User) error {
	r.users = append(r.users, user)
	return nil
}

var userRepository *DBUserRepository

func GetUserRepository() (r UserRepository) {
	if userRepository == nil {
		userRepository = New()
	}
	return userRepository
}
