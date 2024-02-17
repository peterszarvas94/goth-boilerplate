package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/configloader"
)

type config struct {
	PORT             string
	APP_ENV          string
	DB_HOST          string
	DB_PORT          string
	DB_DATABASE      string
	DB_USERNAME      string
	DB_PASSWORD      string
	DB_ROOT_PASSWORD string
}

var App config

func init() {
	var app config

	// Check APP_ENV before other env vars.
	// This is mandatory to set for the app.
	appEnv, ok := os.LookupEnv("APP_ENV")
	if !ok || appEnv == "" {
		fmt.Println("APP_ENV is not set")
		os.Exit(1)
	}

	// Check if we are in a test environment, return so we don't get errors in
	// tests because of missing environment variables.
	if appEnv == "test" {
		return
	}

	err := configloader.Load(&app)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	App = app
}
