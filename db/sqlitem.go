package db

import (
	"github.com/ieshan/zen/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SetupMySQL assigns gorm.DB interface based on config to DBClient
func ConnectSQLite() {
	dbOnce.Do(func() {
		log := logger.GetLogger()
		c, err := gorm.Open("sqlite", ":memory:")
		log.ErrorFatal("Got error when connecting with SQLite in memory database", err)
		dbConn = &DBClient{
			DB: c,
		}
	})
}

// DefaultDB returns the default DBClient currently in Use
func GetSQLiteConn() *DBClient {
	return dbConn
}

