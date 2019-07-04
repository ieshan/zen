package config

import "github.com/spf13/viper"

type LogConfig struct {
	LogLevel string
	Formatter string
}

var logConfig *LogConfig

func GetLog() *LogConfig {
	return logConfig
}

func initLog()  {
	if logConfig != nil {
		return
	}
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.formatter", "json")
	logConfig = &LogConfig{
		LogLevel: viper.GetString("log.level"),
		Formatter: viper.GetString("log.formatter"),
	}
}
