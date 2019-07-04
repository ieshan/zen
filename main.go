package main

import (
	"github.com/ieshan/zen/cmd"
	"github.com/ieshan/zen/config"
	"github.com/ieshan/zen/db"
	"github.com/ieshan/zen/logger"
	"os"
)

func main() {
	config.Init()
	logger.InitLogger(os.Stdout)
	db.ConnectMySQL()
	db.ConnectRedis()
	cmd.Execute()
}
