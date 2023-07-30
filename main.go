package main

import (
	"os"
	"path/filepath"

	"github.com/ivc/dotfiles/cmd"
	"github.com/ivc/dotfiles/pkg/cli"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	app := &cli.App{
		BasePath: filepath.Join(homeDir, ".local", "share", "dotfiles"),
	}
	rootCmd := cmd.NewRootCmd(app)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
