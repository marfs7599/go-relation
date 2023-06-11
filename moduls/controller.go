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

		config.DB.Create(&bpjs)

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

		config.DB.Create(&recipe)

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

		config.DB.Create(&tag)

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
