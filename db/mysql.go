package db

import (
	"fmt"
	"github.com/ieshan/zen/config"
	"github.com/ieshan/zen/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"time"
)

// SetupMySQL assigns gorm.DB interface based on config to MySQLClient
func ConnectMySQL() {
	dbOnce.Do(func() {
		cnf := config.GetDB()
		dbUser := cnf.Username
		dbPass := cnf.Password
		dbHost := cnf.Host
		dbPort := cnf.Port
		dbName := cnf.Name
		dbSource := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			dbUser,
			dbPass,
			dbHost,
			dbPort,
			dbName,
		)
		c, err := gorm.Open("mysql", dbSource)
		log := logger.GetLogger()
		log.ErrorFatal("Got error when connecting with MySQL database", err)

		/*if cnf.MaxLifeTime != 0 {
			c.DB().SetConnMaxLifetime(time.Duration(cnf.MaxLifeTime) * time.Second)
			fmt.Println("SetConnMaxLifetime: ", cnf.MaxLifeTime)
		}
		if cnf.MaxIdleConn != 0 {
			c.DB().SetMaxIdleConns(cnf.MaxIdleConn)
			fmt.Println("SetMaxIdleConns: ", cnf.MaxIdleConn)
		}
		if cnf.MaxOpenConn != 0 {
			c.DB().SetMaxOpenConns(cnf.MaxOpenConn)
			fmt.Println("SetMaxOpenConns: ", cnf.MaxOpenConn)
		}*/

		dbConn = &DBClient{
			DB: c,
		}
	})
}

// DefaultDB returns the default DBClient currently in Use
func GetMySQLConn() *DBClient {
	return dbConn
}
