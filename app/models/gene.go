package models

import (
	"github.com/jinzhu/gorm"
)

type Gene struct {
	gorm.Model
	Id   int
	Name string
	Path string
}
