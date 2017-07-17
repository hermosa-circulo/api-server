package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Gene struct {
	gorm.Model
	Name        string
	Path        string
	Thumb       string
	Views       int
	Description string
	Date        time.Time
}

func (gene Gene) DateString() string {
	return fmt.Sprintf("更新日時：%s", gene.Date.Format("Mon, 02 Jan 2006 15:04:05 MST"))
}
