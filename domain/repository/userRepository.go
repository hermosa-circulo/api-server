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
