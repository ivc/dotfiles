package cmd

import (
	"github.com/ivc/dotfiles/pkg/cli"
	"github.com/spf13/cobra"
)

func newInstallCmd(app *cli.App) *cobra.Command {
	installCmd := &cobra.Command{
		Use:   "install repository...",
		Short: "Install dotfiles from one or more Git repositories.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Install(args...)
		},
	}

	return installCmd
}
