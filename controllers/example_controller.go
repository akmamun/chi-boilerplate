package controllers

import (
	"chi-boilerplate/models"
	"encoding/json"
	"net/http"
)

func CreateExample(w http.ResponseWriter, request *http.Request) {
	example := new(models.Example)
	err := json.NewDecoder(request.Body).Decode(&example)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.SaveExample(&example)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&example)
}

func GetData(w http.ResponseWriter, request *http.Request) {
	var example []models.Example
	models.GetAll(&example)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&example)
}
