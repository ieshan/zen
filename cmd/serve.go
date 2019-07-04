package cmd

import (
	"github.com/ieshan/zen/server"
	"github.com/spf13/cobra"
)

var serverPort uint16

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve HTTP server (Default port: 8000)",
	Run: func(cmd *cobra.Command, args []string) {
		server.ServeHTTP(int(serverPort))
	},
}

func init() {
	serveCmd.Flags().Uint16VarP(&serverPort, "port", "p", 8000, "Set port")
	rootCmd.AddCommand(serveCmd)
}
