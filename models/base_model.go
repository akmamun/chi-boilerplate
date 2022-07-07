package models

import (
	"chi-boilerplate/pkg/database"
	"chi-boilerplate/pkg/logger"
)

func save(value interface{}) {
	err := database.DB.Create(&value).Error
	if err != nil {
		logger.Errorf("cant save", err)
	}
}
