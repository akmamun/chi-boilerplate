package migrations

import (
	"chi-boilerplate/models"
	"chi-boilerplate/pkg/database"
)

func Migrate() {
	var migrationModels = []interface{}{&models.Example{}}
	err := database.GetDB().AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
