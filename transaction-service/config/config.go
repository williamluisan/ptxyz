package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB *DatabaseConfig
}

func Load() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to load viper configuration: " + err.Error())
	}

	return &Config{
		DB: &DatabaseConfig{
			Driver: viper.GetString("DB_DRIVER"),
			Host: viper.GetString("DB_HOST"),
			Port: viper.GetInt("DB_PORT"),
			User: viper.GetString("DB_USER"),
			Pass: viper.GetString("DB_PASSWORD"), 
			Name: viper.GetString("DB_NAME"),
			Charset: viper.GetString("DB_CHARSET"),
		},
	}, nil
}