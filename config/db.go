package config

import "github.com/spf13/viper"

// Database holds the database configuration
type Database struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Name        string
	//MaxLifeTime int // in seconds
	//MaxIdleConn int
	//MaxOpenConn int
}

var dbConfig *Database

func GetDB() *Database {
	return dbConfig
}

func initDB()  {
	if dbConfig != nil {
		return
	}
	viper.SetDefault("database.name", "default")
	viper.SetDefault("database.username", "username")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	dbConfig = &Database{
		Name:        viper.GetString("database.name"),
		Username:    viper.GetString("database.user"),
		Password:    viper.GetString("database.pass"),
		Host:        viper.GetString("database.host"),
		Port:        viper.GetString("database.port"),
		//MaxLifeTime: viper.GetInt("database.max_life_time"),
		//MaxIdleConn: viper.GetInt("database.max_idle_conn"),
		//MaxOpenConn: viper.GetInt("database.max_open_conn"),
	}
}
