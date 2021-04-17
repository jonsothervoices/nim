package main

import (
	"github.com/gorilla/mux"
)

func newAPI() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/play", playPostHandler).Methods("POST")
	r.HandleFunc("/play", playGetHandler).Methods("GET")
	//NOT VALID REST
	r.HandleFunc("/varplay/{coins}/{move}", varPlayGetHandler).Methods("GET")
	return r
}
