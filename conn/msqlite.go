package conn

import (
	"github.com/ieshan/zen/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetInMemorySQLiteConn() *DBClient {
	c, err := gorm.Open("sqlite3", ":memory:")
	log := logger.GetLogger()
	log.ErrorFatal("Failed creating sqlite in-memory", err)
	return &DBClient{
		DB: c,
	}
}
