package git

import (
	"os"
	"os/exec"
)

func (repo *Repo) Push() (err error) {
	// Run `git push origin master`.
	err = os.Chdir(repo.Path)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "push", "origin", "master")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
