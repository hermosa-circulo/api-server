package main

import (
	"fmt"
	"github.com/hermosa-circulo/iga-controller/iga"
	"net/http"
	"os"
)

const (
	defaultPort = ":8080"
)

func getEnv(env, defaultValue string) string {
	e := os.Getenv(env)
	if e == "" {
		return defaultValue
	}
	return e
}

func main() {
	port := getEnv("PORT", defaultPort)
	fmt.Println("starting server with ", port, "...")

	http.HandleFunc("/", iga.Healthcheck)
	http.HandleFunc("/update", iga.Update)
	http.ListenAndServe(port, nil)

}
