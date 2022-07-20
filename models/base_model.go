package models

import (
	"chi-boilerplate/pkg/database"
	"chi-boilerplate/pkg/logger"
)

func Save(value interface{}) interface{} {
	err := database.GetDB().Create(value).Error
	if err != nil {
		logger.Errorf("error, not save data %v", err)
	}
	return err
}

func Get(value interface{}) interface{} {
	err := database.GetDB().Find(value).Error
	return err
}
