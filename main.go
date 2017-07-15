package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Gene struct {
	gorm.Model
	Id    int
	Param int
}

func main() {
	fmt.Println("vim-go")
	db, _ := gorm.Open("mysql", "root:docker@db/api")

	db.AutoMigrate(&Gene{})
	db.Create(&Gene{Id: 1, Param: 1000})

	var gene Gene
	db.First(&gene, 1)
	db.First(&gene, "Id = ?", 1)

	db.Model(&gene).Update("Param", 2000)
	defer db.Close()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
