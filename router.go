package main

import (
	"encoding/json"
	"net/http"
	"iga-controller/domain/models"
)

type AppHandler struct{}

func (a *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func foo(w http.ResponseWriter, r *http.Request) {
	user := models.User{Name: "Alex", Evaluations: []models.Evaluation{}}

	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
