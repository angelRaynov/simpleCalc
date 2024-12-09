package config

import (
	"log"
	"os"
)

type Application struct {
	AppMode string
	APIPort string
}

func New() *Application {
	var config Application

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		log.Fatal("API_PORT not set")
	}
	config.APIPort = apiPort

	appMode := os.Getenv("APP_MODE")
	if appMode == "" {
		appMode = "debug"
	}

	return &config
}
