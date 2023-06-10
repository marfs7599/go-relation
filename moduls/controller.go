package moduls

import (
	"encoding/json"
	"net/http"

	"relation/config"
	"relation/entity"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GetAllPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var patient []entity.Patient

		config.DB.Find(&patient)

		result := Result{Code: 200, Data: patient, Message: "Success collect data patient!"}
		jsonData, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	http.Error(w, "Method not valid!", http.StatusInternalServerError)
}

func InsertPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var patient entity.Patient
		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		config.DB.Create(&patient)

		result := Result{Code: 200, Data: patient, Message: "Success insert data patient!"}
		jsonData, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	http.Error(w, "Method not valid!", http.StatusInternalServerError)
}
