package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading config file.", err)
	}
	initLog()
	initDB()
	initRedis()
}
