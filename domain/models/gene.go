package models

import (
	"github.com/jinzhu/gorm"
)

type Gene struct {
	gorm.Model
	Name        string
	Path        string
	Evaluations []Evaluation
}
