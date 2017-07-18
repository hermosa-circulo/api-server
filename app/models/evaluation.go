package models

import (
	"github.com/jinzhu/gorm"
)

type Evaluation struct {
	gorm.Model
	Id     int
	UserId int
	Value  int
	GeneId int
}
