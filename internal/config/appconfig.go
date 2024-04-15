package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var EnvConfigs *envConfigs

type envConfigs struct {
	SECRET_KEY            string `mapstructure:"SECRET_KEY"`
	PRIVATE_SIGN_KEY_PATH string `mapstructure:"PRIVATE_SIGN_KEY_PATH"`
	PUBLIC_SIGN_KEY_PATH  string `mapstructure:"PUBLIC_SIGN_KEY_PATH"`
	PRIVATE_SIGN_KEY      []byte `mapstructure:"PRIVATE_SIGN_KEY"`
	PUBLIC_SIGN_KEY       []byte `mapstructure:"PUBLIC_SIGN_KEY"`
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
	err := setPrivateSignKey(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := setPublicSignKey(config); err != nil {
		log.Fatal(err)
	}
	return
}

func setPrivateSignKey(config *envConfigs) error {
	privateKey, err := os.ReadFile("C:\\Users\\bryanspc\\GolandProjects\\login_app\\private.key.pem")
	if err != nil {
		return err
	}
	config.PRIVATE_SIGN_KEY = privateKey
	return nil

}

func setPublicSignKey(config *envConfigs) error {
	publicKey, err := os.ReadFile("C:\\Users\\bryanspc\\GolandProjects\\login_app\\public.key.pem")
	if err != nil {
		return err
	}
	config.PUBLIC_SIGN_KEY = publicKey
	return nil
}
