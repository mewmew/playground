// TODO(u): remove hardcoded username to make the tool usable for more general
// purposes.

// gitsync keeps forked git repositories in sync with their parents.
package main

import (
	"fmt"
	"log"

	"github.com/mewfork/go-github/github"
)

func main() {
	err := gitsync()
	if err != nil {
		log.Fatalln(err)
	}
}

func gitsync() (err error) {
	c := github.NewClient(nil)

	repos, _, err := c.Repositories.List("mewpkg", nil)
	if err != nil {
		return err
	}
	fmt.Println("BASE_DIR = $PWD")
	fmt.Println("eval `ssh-agent`")
	fmt.Println("ssh-add ~/.ssh/id_rsa_mewmew")

	for _, r := range repos {
		repo, _, err := c.Repositories.Get("mewpkg", *r.Name)
		if err != nil {
			return err
		}
		if *repo.Fork {
			fmt.Println("cd $BASE_DIR")
			fmt.Println("git clone", *repo.CloneURL)
			fmt.Printf("cd $BASE_DIR/%s\n", *repo.Name)
			fmt.Println("git pull", *repo.Parent.CloneURL)
			fmt.Println("git push -u origin master")
		}
	}

	return nil
}
