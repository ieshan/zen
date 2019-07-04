package cmd

import (
	"github.com/ieshan/zen/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zen",
	Short: "Zen is a Echo and Vue based Go project boilerplate",
}

func Execute() {
	log := logger.GetLogger()
	log.ErrorFatal("Root command execution error.", rootCmd.Execute())
}
