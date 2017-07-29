package models

import (
	"github.com/jinzhu/gorm"
)

type Evaluation struct {
	gorm.Model
	Value  int
	UserId int
	User   User
}
