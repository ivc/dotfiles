package cmd

import (
	"github.com/ivc/dotfiles/pkg/cli"
	"github.com/spf13/cobra"
)

// NewRootCmd builds a new root cobra.Command for the given cli.App.
func NewRootCmd(app *cli.App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dotfiles",
		Short: "A command-line tool for managing and synchronizing dotfiles via Git repositories.",
	}

	return rootCmd
}
