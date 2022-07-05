package controllers

import (
	"chi-boilerplate/models"
	"chi-boilerplate/pkg/logger"
	"encoding/json"
	"net/http"
)

func (base *BaseController) CreateExample(w http.ResponseWriter, request *http.Request) {
	example := new(models.Example)
	err := json.NewDecoder(request.Body).Decode(&example)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = base.DB.Create(&example).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&example)
}

func (base *BaseController) GetData(w http.ResponseWriter, request *http.Request) {
	var example []models.Example
	base.DB.Find(&example)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&example)
}
