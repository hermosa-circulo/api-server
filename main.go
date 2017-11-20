package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	defaultPort = ":8080"
)

var logger = func(method, uri, name string, start time.Time) {
	log.Printf("\"method\":%q  \"uri\":%q    \"name\":%q   \"time\":%q", method, uri, name, time.Since(start))
}

func getEnv(env, defaultValue string) string {
	e := os.Getenv(env)
	if e == "" {
		return defaultValue
	}
	return e
}

func main() {

	port := getEnv("PORT", defaultPort)
	router := httprouter.New()

	fmt.Println("starting server with ", port, "...")

	router.GET("/", Logging(Index, "index"))
	router.GET("/health", Logging(Healthcheck, "healhcheck"))
	router.POST("/update", Logging(Update, "update"))

	log.Fatal(http.ListenAndServe(port, router))

}

func Logging(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		logger(r.Method, r.URL.Path, name, start)
	}
}
