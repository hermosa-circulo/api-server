package controllers

import (
	"iga-controller/app/models"

	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(models.InitDB)
}
