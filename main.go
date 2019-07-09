package main

import (
	"github.com/ieshan/zen/cmd"
	"github.com/ieshan/zen/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
