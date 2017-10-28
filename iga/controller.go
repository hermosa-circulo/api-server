package iga

import (
	"encoding/json"
	"net/http"
)

type AppHandler struct{}

func (a *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func Update(w http.ResponseWriter, r *http.Request) {

	js, err := json.Marshal("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
