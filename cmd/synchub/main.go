package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

func main() {
	var (
		// skipKey specifies whether to skip using a GitHub SSH key.
		skipKey bool
	)
	flag.BoolVar(&skipKey, "skipKey", false, "Skip using GitHub SSH key.")
	flag.Parse()
	if err := synchub(skipKey); err != nil {
		log.Fatal(err)
	}
}

func synchub(skipKey bool) error {
	buf, err := ioutil.ReadFile("access_token.txt")
	if err != nil {
		return errors.WithStack(err)
	}
	token := string(bytes.TrimSpace(buf))
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	fmt.Println("#!/bin/bash")
	fmt.Println()

	// Generate GitHub sync script.
	if !skipKey {
		fmt.Println("# Add GitHub SSH key.")
		fmt.Println("eval `ssh-agent`")
		fmt.Println("ssh-add")
		fmt.Println()
	}
	user := "mewmew"
	repos, _, err := client.Repositories.List(ctx, user, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := genSyncScript(ctx, client, repos, skipKey); err != nil {
		return errors.WithStack(err)
	}
	orgs, _, err := client.Organizations.List(ctx, user, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, org := range orgs {
		orgName := org.GetLogin()
		opt := &github.RepositoryListByOrgOptions{
			ListOptions: github.ListOptions{PerPage: 10},
		}
		var repos []*github.Repository
		for {
			rs, resp, err := client.Repositories.ListByOrg(ctx, orgName, opt)
			if err != nil {
				return errors.WithStack(err)
			}
			repos = append(repos, rs...)
			if resp.NextPage == 0 {
				break
			}
			opt.ListOptions.Page = resp.NextPage
		}
		if err := genSyncScript(ctx, client, repos, skipKey); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func genSyncScript(ctx context.Context, client *github.Client, repos []*github.Repository, skipKey bool) error {
	for _, repo := range repos {
		t, err := template.New("script").Parse(script)
		if err != nil {
			return errors.WithStack(err)
		}
		owner := repo.Owner.GetLogin()
		repoName := repo.GetName()
		repoPath := fmt.Sprintf("github.com/%s/%s", owner, repoName)
		cloneURL := repo.GetCloneURL()
		sshURL := repo.GetSSHURL()
		defaultBranch := repo.GetDefaultBranch()
		data := map[string]interface{}{
			"Owner":         owner,
			"RepoName":      repoName,
			"RepoPath":      repoPath,
			"CloneURL":      cloneURL,
			"SSHURL":        sshURL,
			"DefaultBranch": defaultBranch,
			"UseSSHKey":     !skipKey,
		}
		if repo.GetFork() {
			data["Fork"] = true
			repository, _, err := client.Repositories.Get(ctx, owner, repoName)
			if err != nil {
				return errors.WithStack(err)
			}
			data["UpstreamURL"] = repository.Parent.GetCloneURL()
		}
		if err := t.Execute(os.Stdout, data); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

const script = `# Clone {{ .RepoPath }}
echo "Cloning {{ .RepoPath }}"
OrgPath="github.com/{{ .Owner }}"
RepoPath="{{ .RepoPath }}"
mkdir -p ${OrgPath}
git -C ${OrgPath} clone {{ .CloneURL }}
git -C ${RepoPath} remote set-url origin {{ .SSHURL }}
{{- if .Fork }}
# Sync with upstream.
git -C ${RepoPath} remote add upstream {{ .UpstreamURL }}
git -C ${RepoPath} pull upstream {{ .DefaultBranch }}
	{{- if .UseSSHKey }}
git -C ${RepoPath} push -u origin {{ .DefaultBranch }}
	{{- end }}
{{- end }}

`
