package main

import (
	"fmt"
	"github.com/hermosa-circulo/iga-controller/iga"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
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

	http.HandleFunc("/", iga.Index)
	http.HandleFunc("/health", iga.Healthcheck)
	http.HandleFunc("/update", iga.Update)
	log.Fatal(http.ListenAndServe(port, nil))

}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
