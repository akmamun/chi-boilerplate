package models

import (
	"chi-boilerplate/pkg/database"
	"chi-boilerplate/pkg/logger"
	"time"
)

type Example struct {
	Id        int        `json:"id"`
	Data      string     `json:"data" binding:"required"`
	Data2     string     `json:"data2" binding:"required,string, omitempty"`
	Data3     string     `json:"data3" binding:"required,string, omitempty"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

// TableName is Database TableName of this model
func (Example) TableName() string {
	return "examples"
}

//func (e *Example) TableName() string {
//	return "examples"
//}
func (e *Example) AutoMigrate() {
	err := database.GetDB().AutoMigrate(e).Error
	if err != nil {
		logger.Errorf("error")
	}
}

func SaveExample(data interface{}) {
	Save(data)
}
