// gitsync keeps forked git repositories in sync with their parents. It does so
// by locating the repositories of provided usernames and organizations. Then it
// creates a shell script which will clone all repository forks, pull changes
// from their parens and push those changes to the forked repository.
//
// Usage:
//    gitsync USER...
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "gitsync USER...")
}

func main() {
	var (
		// GitHub access token.
		accessToken string
	)
	flag.StringVar(&accessToken, "token", "", "GitHub access token")
	flag.Parse()
	for _, username := range flag.Args() {
		err := gitsync(username, accessToken)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// gitsync locates the repositories of the provided username or organization. It
// creates a shell script which will clone all repository forks, pull changes
// from their parens and push those changes to the forked repository.
func gitsync(username, accessToken string) (err error) {
	c := github.NewClient(nil)
	// Use GitHub access token.
	if len(accessToken) > 0 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		c = github.NewClient(tc)
	}

	repos, _, err := c.Repositories.List(context.TODO(), username, nil)
	if err != nil {
		return err
	}
	fmt.Println("export BASE_DIR=$PWD")
	fmt.Println("eval `ssh-agent`")
	fmt.Println("ssh-add ~/.ssh/id_rsa_mewmew")
	for i, r := range repos {
		log.Printf("   %d/%d", i+1, len(repos))
		repo, _, err := c.Repositories.Get(context.TODO(), username, *r.Name)
		if err != nil {
			return err
		}
		if *repo.Fork {
			gitCloneURL := getGitCloneURL(*repo.CloneURL)
			fmt.Println("cd $BASE_DIR")
			fmt.Println("git clone", gitCloneURL)
			fmt.Printf("cd $BASE_DIR/%s\n", *repo.Name)
			fmt.Println("git pull", *repo.Parent.CloneURL)
			fmt.Println("git push -u origin master")
		}
	}

	return nil
}

func getGitCloneURL(cloneURL string) string {
	return strings.Replace(cloneURL, "https://github.com/", "git@github.com:", -1)
}
