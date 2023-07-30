package git

import (
	"os/exec"
	"strings"
)

// Remote is an alias for a string representing a remote Git repository URL.
type Remote = string

// Name returns the trailing part of the repo after final '/' separator.
func Name(repo Remote) string {
	idx := strings.LastIndexByte(repo, '/')
	return repo[idx+1:]
}

// Clone returns a new 'git clone <repo> <name>' Cmd.
func Clone(repo Remote, name string) *exec.Cmd {
	return exec.Command("git", "clone", repo, name)
}
