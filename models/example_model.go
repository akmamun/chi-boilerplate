package models

import (
	"time"
)

type Example struct {
	Id        int        `json:"id"`
	Data      string     `json:"data" binding:"required"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

// TableName is Database TableName of this model
func (Example) TableName() string {
	return "examples"
}
func SaveExample(data interface{}) {
	Save(data)
}
func GetAll(value interface{}) interface{} {
	err := Get(value)
	return err
}
