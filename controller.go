package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	js, err := json.Marshal("")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	genes := GetGeneService().GetAllGene()

	js, err := json.Marshal(genes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
