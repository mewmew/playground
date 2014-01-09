package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mewmew/playground/cmd/gitstat/git"
)

func main() {
	flag.Parse()
	for _, dir := range flag.Args() {
		err := gitstat(dir)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func gitstat(dir string) (err error) {
	// Check repos.
	repos, err := git.Check(dir)
	if err != nil {
		return err
	}
	var aheadRepos, stagedRepos, unstagedRepos, untrackedRepos []*git.Repo
	for _, repo := range repos {
		if repo.IsAhead() {
			aheadRepos = append(aheadRepos, repo)
		}
		if repo.HasStaged() {
			stagedRepos = append(stagedRepos, repo)
		}
		if repo.HasUnstaged() {
			unstagedRepos = append(unstagedRepos, repo)
		}
		if repo.HasUntracked() {
			untrackedRepos = append(untrackedRepos, repo)
		}
	}

	// === [ repo output ] ======================================================

	// Ahead.
	if len(aheadRepos) > 0 {
		fmt.Println("--- [ ahead ] ---")
		for _, repo := range aheadRepos {
			fmt.Println("repo:", repo.Base)
		}
		fmt.Println()
	}

	// Staged changes.
	if len(stagedRepos) > 0 {
		fmt.Println("--- [ staged changes ] ---")
		for _, repo := range stagedRepos {
			fmt.Println("repo:", repo.Base)
		}
		fmt.Println()
	}

	// Unstaged changes.
	if len(unstagedRepos) > 0 {
		fmt.Println("--- [ unstaged changes ] ---")
		for _, repo := range unstagedRepos {
			fmt.Println("repo:", repo.Base)
		}
		fmt.Println()
	}

	// Untracked files.
	if len(untrackedRepos) > 0 {
		fmt.Println("--- [ untracked files ] ---")
		for _, repo := range untrackedRepos {
			fmt.Println("repo:", repo.Base)
		}
		fmt.Println()
	}

	// === [ repo work ] ========================================================

	// Ahead.
	if len(aheadRepos) > 0 {
		fmt.Println("=== [ ahead ] ===")
		fmt.Println()
		for _, repo := range aheadRepos {
			err = repo.Push()
			if err != nil {
				return err
			}
		}
	}

	// Staged changes.
	if len(stagedRepos) > 0 {
		fmt.Println("=== [ staged changes ] ===")
		fmt.Println()
		for _, repo := range stagedRepos {
			fmt.Printf("--- [ %s ] ---\n", repo.Base)
			err = repo.Status()
			if err != nil {
				return err
			}
			fmt.Println()
		}
	}

	// Unstaged changes.
	if len(unstagedRepos) > 0 {
		fmt.Println("=== [ unstaged changes ] ===")
		fmt.Println()
		for _, repo := range unstagedRepos {
			fmt.Printf("--- [ %s ] ---\n", repo.Base)
			err = repo.Status()
			if err != nil {
				return err
			}
			fmt.Println()
		}
	}

	return nil
}
