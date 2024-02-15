package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/configloader"
)

type config struct {
	PORT string
	APP_ENV string
	DB_HOST string
	DB_PORT string
	DB_DATABASE string
	DB_USERNAME string
	DB_PASSWORD string
	DB_ROOT_PASSWORD string
}

var App config

func init() {
	err := configloader.Load(&App)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
