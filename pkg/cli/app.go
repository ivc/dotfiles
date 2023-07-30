package cli

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ivc/dotfiles/pkg/git"
)

// App provides CLI implementation.
type App struct {
	BasePath string
}

// Install downloads repositories and establishes links in $HOME directory.
func (app *App) Install(repositories ...git.Remote) error {
	var errs []error
	for _, repo := range repositories {
		err := app.install(git.Name(repo), repo)
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

func (app *App) gitCmd(dir string, cmd *exec.Cmd) *exec.Cmd {
	cmd.Dir = filepath.Join(app.BasePath, dir)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func (app *App) install(repoName string, repo git.Remote) error {
	cloneCmd := app.gitCmd("", git.Clone(repo, repoName))
	if err := os.MkdirAll(cloneCmd.Dir, 0700); err != nil {
		return err
	}
	return cloneCmd.Run()
}
