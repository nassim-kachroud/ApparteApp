package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	Port              string
	DbHost            string
	DbPort            int
	DbName            string
	DbUser            string
	DbPassword        string
	AirbnbAPIKey      string
	AirbnbAccessToken string
)

//Get...
func Get() {
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("WEB_REQUEST_TIMEOUT", 10000)
	viper.SetDefault("WEB_SERVER_TIMEOUT", 10000)
	viper.SetDefault("AirbnbAPIKey", "d306zoyjsyarp7ifhu67rjxn52tv0t20")
	viper.SetDefault("AirbnbAPIKey", "")

	if os.Getenv("ENVIRONMENT") == "DEV" {
		_, dirname, _, _ := runtime.Caller(0)
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(filepath.Dir(dirname))
		viper.ReadInConfig()
	} else {
		viper.AutomaticEnv()
	}

	//Assign env variables value to global variables
	Port = viper.GetString("APP_PORT")
	DbHost = viper.GetString("DB_HOST")
	DbPort = viper.GetInt("DB_PORT")
	DbName = viper.GetString("DB_NAME")
	DbUser = viper.GetString("DB_USER")
	DbPassword = viper.GetString("DB_PASSWORD")
	AirbnbAPIKey = viper.GetString("AIRBNB_API_KEY")
	AirbnbAccessToken = viper.GetString("AIRBNB_ACCESS_TOKEN")
	// WebRequestTimeout = viper.GetInt64("WEB_REQUEST_TIMEOUT")
	// WebServerTimeout = time.Duration(viper.GetInt64("WEB_SERVER_TIMEOUT")) * time.Millisecond
	// DbRequestTimeout = viper.GetInt("DB_REQUEST_TIMEOUT")
}
