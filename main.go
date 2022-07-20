package main

import (
	"chi-boilerplate/migrations"
	"chi-boilerplate/pkg/config"
	"chi-boilerplate/pkg/database"
	"chi-boilerplate/pkg/logger"
	"chi-boilerplate/routers"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {

	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	if err := database.SetupConnection(); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}
	//TODO changes in later
	migrations.Migrate()

	router := routers.SetupRoute()

	logger.Fatalf("%v", http.ListenAndServe(config.ServerConfig(), router))

}
