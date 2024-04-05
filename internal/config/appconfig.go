package config

import (
	"github.com/spf13/viper"
	"log"
)

var EnvConfigs *envConfigs

type envConfigs struct {
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
}

func InitEnvConfigs() {
	EnvConfigs = loadEnvConfigs()
}

func loadEnvConfigs() (config *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
