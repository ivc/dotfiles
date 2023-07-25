package main

import (
	"os"

	"github.com/ivc/dotfiles/cmd"
	"github.com/ivc/dotfiles/pkg/cli"
)

func main() {
	app := &cli.App{}
	rootCmd := cmd.NewRootCmd(app)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
