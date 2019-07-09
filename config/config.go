package config

import (
	"github.com/ieshan/zen/logger"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	log := logger.GetLogger()
	log.ErrorFatal("Error loading config file.", viper.ReadInConfig())
	initDB()
	initRedis()
}
