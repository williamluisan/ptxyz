package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {}

func Load() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to load viper configuration: " + err.Error())
	}

	return &Config{}, nil
}