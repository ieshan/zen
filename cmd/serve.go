package cmd

import (
	"github.com/ieshan/zen/instance"
	"github.com/ieshan/zen/logger"
	"github.com/ieshan/zen/route"
	"github.com/spf13/cobra"
	"strconv"
)

var serverPort uint16

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve HTTP server (Default port: 8000)",
	Run: func(cmd *cobra.Command, args []string) {
		e := instance.GetEcho()
		db := instance.GetDB()
		route.Initialize(e, db)
		//route.DummyRoutes(e.Group("/user"))
		log := logger.GetLogger()
		log.ErrorFatal("ServeHTTP error.", e.Start(":" + strconv.Itoa(int(serverPort))))
	},
}

func init() {
	serveCmd.Flags().Uint16VarP(&serverPort, "port", "p", 8000, "Set port")
	rootCmd.AddCommand(serveCmd)
}
