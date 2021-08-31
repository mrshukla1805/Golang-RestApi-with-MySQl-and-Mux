package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func start() {
	r := mux.NewRouter()

	r.HandleFunc("/projects", GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", GetProject).Methods("GET")
	r.HandleFunc("/projects", CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", DeleteProject).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	start()
	ApiCall()

}
