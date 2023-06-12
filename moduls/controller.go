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
		var patients []entity.Patient

		config.DB.Preload("Bpjs").Preload("Recipe").Find(&patients)

		result := Result{Code: 200, Data: patients, Message: "Success collect data patients!"}
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

		// ======== Manual Validation =======
		if patient.Nik == 0 {
			http.Error(w, "ID Pasien is required", 400)
			return
		}

		if patient.Name == "" {
			http.Error(w, "Name is required", 400)
			return
		}

		if patient.Address == "" {
			http.Error(w, "Address is required", 400)
			return
		}

		if patient.Phonenumber == "" {
			http.Error(w, "Phonenumber is required", 400)
			return
		}
		// ==== End Validation ====

		if err := config.DB.Create(&patient).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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

func GetAllBpjs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var bpjs []entity.Bpjs

		config.DB.Preload("Patient").Find(&bpjs)

		result := Result{Code: 200, Data: bpjs, Message: "Success collect data bpjs!"}
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

func InsertBpjs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var bpjs entity.Bpjs
		err := json.NewDecoder(r.Body).Decode(&bpjs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ======== Manual Validation =======
		if bpjs.NoCard == "" {
			http.Error(w, "No Card is required", 400)
			return
		}

		if bpjs.PatientId == 0 {
			http.Error(w, "ID Pasien is required", 400)
			return
		}
		// ==== End Validation ====

		if err := config.DB.Create(&bpjs).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := Result{Code: 200, Data: bpjs, Message: "Success insert data bpjs!"}
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

func GetAllRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var recipes []entity.RecipeResponseTag

		config.DB.Preload("Patient").Preload("Tags").Find(&recipes)

		result := Result{Code: 200, Data: recipes, Message: "Success collect data recipes!"}
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

func InsertRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var recipe entity.Recipe
		err := json.NewDecoder(r.Body).Decode(&recipe)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ======== Manual Validation =======
		if recipe.Name == "" {
			http.Error(w, "Name is required", 400)
			return
		}

		if recipe.Dose == "" {
			http.Error(w, "Dose is required", 400)
			return
		}

		if recipe.PatientId == 0 {
			http.Error(w, "ID Pasien is required", 400)
			return
		}

		if len(recipe.TagsId) == 0 {
			http.Error(w, "ID Pasien is required", 400)
			return
		}
		// ==== End Validation ====

		if err := config.DB.Create(&recipe).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(recipe.TagsId) > 0 {
			for _, TagID := range recipe.TagsId {
				recipeTag := new(entity.RecipeTag)
				recipeTag.RecipeId = recipe.Id
				recipeTag.TagId = TagID
				config.DB.Create(&recipeTag)
			}
		}

		result := Result{Code: 200, Data: recipe, Message: "Success insert data recipe!"}
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

func GetAllTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var tags []entity.TagResponse

		config.DB.Preload("Recipe").Find(&tags)

		result := Result{Code: 200, Data: tags, Message: "Success collect data tags!"}
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

func InsertTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var tag entity.Tag
		err := json.NewDecoder(r.Body).Decode(&tag)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ======== Manual Validation =======
		if tag.Name == "" {
			http.Error(w, "Name is required", 400)
			return
		}
		// ==== End Validation ====

		if err := config.DB.Create(&tag).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := Result{Code: 200, Data: tag, Message: "Success insert data tag!"}
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
