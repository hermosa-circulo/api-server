package models

import (
	"github.com/jinzhu/gorm"
)

type Gene struct {
	gorm.Model
	Name        string
	Evaluations []Evaluation
	WheelRadius int
	Begin       int
	PointNum    int
	BreastWide  int
}
