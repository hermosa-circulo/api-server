package main

import (
	"fmt"
	"net/http"
)

const port = ":1312"

func main() {
	fmt.Println("starting server with ", port, "...")
	http.HandleFunc("/", health)
	http.HandleFunc("/user", foo)
	http.ListenAndServe(port, nil)
}
