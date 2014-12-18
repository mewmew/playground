package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Repo struct {
	Base string
	Path string
	Stat
}

func Check(dir string) (repos []*Repo, err error) {
	// Locate git repos.
	walk := func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() && fi.Name() == ".git" {
			// Append located git repo.
			repo := new(Repo)
			repo.Base = filepath.Clean(path[:len(path)-len(".git")])
			repo.Path, err = filepath.Abs(repo.Base)
			if err != nil {
				return err
			}
			repos = append(repos, repo)
			return filepath.SkipDir
		}
		return nil
	}
	err = filepath.Walk(dir, walk)
	if err != nil {
		return nil, err
	}

	// Get git repo status.
	for _, repo := range repos {
		err = repo.ParseStatus()
		if err != nil {
			return nil, err
		}
	}

	return repos, nil
}

const (
	Clean     Stat = 1 << iota // "nothing to commit, working directory clean"
	Untracked                  // "# Untracked files:"
	Unstaged                   // "# Changes not staged for commit"
	Staged                     // "# Changes to be committed:"
	Ahead                      // "Your branch is ahead of 'origin/master' by 1 commit"
)

type Stat int

func (stat Stat) IsClean() bool {
	return stat&Clean != 0
}

func (stat Stat) HasUntracked() bool {
	return stat&Untracked != 0
}

func (stat Stat) HasUnstaged() bool {
	return stat&Unstaged != 0
}

func (stat Stat) HasStaged() bool {
	return stat&Staged != 0
}

func (stat Stat) IsAhead() bool {
	return stat&Ahead != 0
}

var bytesClean = []byte("nothing to commit, working directory clean")
var bytesUntracked = []byte("Untracked files:")
var bytesUnstaged = []byte("Changes not staged for commit:")
var bytesStaged = []byte("Changes to be committed:")
var bytesAhead = []byte("Your branch is ahead of")

func (repo *Repo) ParseStatus() (err error) {
	// Get output from `git status` command.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(cwd)
	err = os.Chdir(repo.Path)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "status")
	buf, err := cmd.Output()
	if err != nil {
		return err
	}

	// Parse output.
	if bytes.Contains(buf, bytesClean) {
		repo.Stat |= Clean
	}
	if bytes.Contains(buf, bytesUntracked) {
		repo.Stat |= Untracked
	}
	if bytes.Contains(buf, bytesUnstaged) {
		repo.Stat |= Unstaged
	}
	if bytes.Contains(buf, bytesStaged) {
		repo.Stat |= Staged
	}
	if bytes.Contains(buf, bytesAhead) {
		repo.Stat |= Ahead
	}

	return nil
}

func (repo *Repo) String() string {
	var m = map[Stat]string{
		Clean:     "clean",
		Untracked: "untracked files",
		Unstaged:  "unstaged changes",
		Staged:    "staged changes",
		Ahead:     "ahead",
	}
	var ss []string
	for mask := Clean; mask <= Ahead; mask <<= 1 {
		if repo.Stat&mask != 0 {
			ss = append(ss, m[mask])
		}
	}
	return fmt.Sprintf("%s (%s)", repo.Path, strings.Join(ss, " | "))
}

func (repo *Repo) Status() (err error) {
	// Run `git push origin master`.
	err = os.Chdir(repo.Path)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "status")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
