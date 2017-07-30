package main

import (
	"fmt"
	"iga-controller/domain/models"
	"net/http"
)

const port = ":1312"

func main() {
	fmt.Println("starting server with ", port, "...")
	models.InitDB()

	http.HandleFunc("/", health)
	http.HandleFunc("/user", foo)
	http.ListenAndServe(port, nil)
}
