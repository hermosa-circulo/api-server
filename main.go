package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Gene struct {
	gorm.Model
	Id    int
	Param int
}

func main() {
	fmt.Println("vim-go")
	db, _ := gorm.Open("mysql", "root:api@localhost/api?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
