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

	"github.com/google/go-github/v28/github"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: gitsync [OPTION]... USER...")
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
			log.Fatalf("%+v", err)
		}
	}
}

// gitsync locates the repositories of the provided username or organization. It
// creates a shell script which will clone all repository forks, pull changes
// from their parens and push those changes to the forked repository.
func gitsync(username, accessToken string) error {
	// Use GitHub access token.
	c := github.NewClient(nil)
	if len(accessToken) > 0 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		c = github.NewClient(tc)
	}
	// Get list of all repos using pagination.
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}
	var allRepos []*github.Repository
	for i := 0; ; i++ {
		log.Printf("getting repository list (page %d)", i+1)
		repos, resp, err := c.Repositories.List(context.TODO(), username, opt)
		if err != nil {
			return errors.WithStack(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	// Get parent repo of all repos.
	for i, r := range allRepos {
		log.Printf("getting additional repository information (repo %d/%d)", i+1, len(allRepos))
		// Get repo.Parent.
		repo, _, err := c.Repositories.Get(context.TODO(), username, *r.Name)
		if err != nil {
			log.Printf("%+v", errors.WithStack(err))
			break
			//return errors.WithStack(err)
		}
		allRepos[i] = repo
	}
	// Print sync script.
	fmt.Println("export BASE_DIR=$PWD")
	fmt.Println("eval `ssh-agent`")
	fmt.Println("ssh-add ~/.ssh/id_rsa")
	for _, repo := range allRepos {
		if *repo.Fork {
			gitCloneURL := getGitCloneURL(*repo.CloneURL)
			fmt.Println("cd $BASE_DIR")
			fmt.Println("git clone", gitCloneURL)
			fmt.Printf("cd $BASE_DIR/%s\n", *repo.Name)
			if repo.Parent != nil {
				fmt.Println("git pull", *repo.Parent.CloneURL)
				fmt.Println("git push -u origin master")
			} else {
				fmt.Println("# unknown parent repo (GitHub rate limit hit)")
			}
		}
	}
	return nil
}

// getGitCloneURL replaces https with git scheme in the Git clone URL.
func getGitCloneURL(cloneURL string) string {
	return strings.ReplaceAll(cloneURL, "https://github.com/", "git@github.com:")
}
