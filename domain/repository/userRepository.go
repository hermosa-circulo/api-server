package repository

import (
	"iga-controller/app/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUsers() []*models.User
	GetUserById(id string) (*models.User, error)
	SaveUser(user *models.User) error
}

type DBUserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *DBUserRepository {
	return &DBUserRepository{
		db: *models.Db
		},
	}
}

func (r *DBUserRepository) GetUsers() []*models.User {
	var users []Users
	db.Find(&users)
	return users
}

func (r *DBUserRepository) GetUserById(id int) (*models.User, error) {

	var user User
	db.First(&user, id)
	return user, nil
	//return nil, errors.New("user not found")
}

func (r *DBUserRepository) SaveUser(user *models.User) error {
	db.Create(&user)
	return nil
}

var userRepository *DBUserRepository

func GetUserRepository() (r UserRepository) {
	if userRepository == nil {
		userRepository = NewUserRepository()
	}
	return userRepository
}
