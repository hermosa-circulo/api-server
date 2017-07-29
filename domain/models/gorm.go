package models

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

func InitDB() {
	var err error
	var dsn = os.Getenv("DATABASE_USER") +
		":" + os.Getenv("DATABASE_PASSWORD") +
		"@" + os.Getenv("DATABASE_HOSTNAME") +
		"/" + os.Getenv("DATABASE_NAME") +
		"?parseTime=true&loc=Asia%2FTokyo"

	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db.LogMode(true)
	autoMigrate()

	log.Println("Connected to database.")
}

func autoMigrate() {
	Db.AutoMigrate(&Gene{})
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Evaluation{})
}
