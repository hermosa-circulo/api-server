package iga

import (
	"encoding/json"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
	w.Write([]byte("ok"))
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

func Index(w http.ResponseWriter, r *http.Request) {

	genes := GetGeneService().GetAllGene()

	js, err := json.Marshal(genes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
