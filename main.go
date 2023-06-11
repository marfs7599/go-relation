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
	r.HandleFunc("/recipes", moduls.GetAllRecipe).Methods("GET")
	r.HandleFunc("/recipes", moduls.InsertRecipe).Methods("POST")
	r.HandleFunc("/tags", moduls.GetAllTag).Methods("GET")
	r.HandleFunc("/tags", moduls.InsertTag).Methods("POST")

	fmt.Println("server start at localhost:9000")
	http.ListenAndServe(":9000", r)
}
