package main

import (
	"fmt"
	"net/http"

	"relation/config"
	"relation/moduls"

	"github.com/gorilla/mux"
)

func main() {
	config.Connect()
	config.Migration()

	r := mux.NewRouter()
	r.HandleFunc("/patients", moduls.GetAllPatient).Methods("GET")
	r.HandleFunc("/patients", moduls.InsertPatient).Methods("POST")
	r.HandleFunc("/bpjs", moduls.GetAllBpjs).Methods("GET")
	r.HandleFunc("/bpjs", moduls.InsertBpjs).Methods("POST")

	fmt.Println("server start at localhost:9000")
	http.ListenAndServe(":9000", r)
}
