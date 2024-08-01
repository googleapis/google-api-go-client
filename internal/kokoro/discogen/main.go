// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v59/github"
	"golang.org/x/oauth2"
)

const (
	branchName     = "discogen"
	commitTitle    = "feat(all): auto-regenerate discovery clients"
	owner          = "googleapis"
	repo           = "google-api-go-client"
	remoteCacheEnv = "REMOTE_CACHE_DIR"
)

func main() {
	ctx := context.Background()

	githubAccessToken := flag.String("github-access-token", os.Getenv("GITHUB_ACCESS_TOKEN"), "The token used to open pull requests. Required.")
	githubUsername := flag.String("github-username", os.Getenv("GITHUB_USERNAME"), "The GitHub user name for the author. Required.")
	githubName := flag.String("github-name", os.Getenv("GITHUB_NAME"), "The name of the author for git commits. Required.")
	githubEmail := flag.String("github-email", os.Getenv("GITHUB_EMAIL"), "The email address of the author. Required.")
	discoDir := flag.String("discovery-dir", os.Getenv("DISCOVERY_DIR"), "Directory where sources of googleapis/google-api-go-client resides. Required.")
	dryRun := flag.Bool("dry-run", false, "Dry run will not commit changes or open a pull request.")

	flag.Parse()

	if *githubAccessToken == "" || *githubUsername == "" || *githubName == "" || *githubEmail == "" || *discoDir == "" {
		log.Fatal("all required flags not set")
	}

	if err := setGitCreds(*githubName, *githubEmail, *githubUsername, *githubAccessToken); err != nil {
		log.Fatalf("unable to set git credentials: %v", err)
	}

	discoSpecDir, err := cloneDiscoArtifactManager()
	if err != nil {
		log.Fatal(err)
	}

	if prIsOpen, err := isPROpen(ctx, *githubAccessToken, *githubUsername); err != nil || prIsOpen {
		if err != nil {
			log.Fatalf("unable to check PR status: %v", err)
		}
		log.Println("a regen PR is already open, nothing to do here")
		os.Exit(0)
	}

	if err := generate(*discoDir, discoSpecDir); err != nil {
		log.Fatalf("unable to generate discovery clients: %v", err)
	}

	if hasChanges, err := hasChanges(*discoDir); err != nil || !hasChanges {
		if err != nil {
			log.Fatalf("unable to check git status: %v", err)
		}
		log.Println("no local changes, exiting")
		os.Exit(0)
	}

	if err := makePR(ctx, *githubAccessToken, *discoDir, *dryRun); err != nil {
		log.Fatalf("unable to make regen PR: %v", err)
	}
}

// setGitCreds configures credentials for GitHub.
func setGitCreds(githubName, githubEmail, githubUsername, accessToken string) error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	gitCredentials := []byte(fmt.Sprintf("https://%s:%s@github.com", githubUsername, accessToken))
	if err := ioutil.WriteFile(path.Join(u.HomeDir, ".git-credentials"), gitCredentials, 0644); err != nil {
		return err
	}
	c := exec.Command("git", "config", "--global", "user.name", githubName)
	c.Env = []string{
		fmt.Sprintf("PATH=%s", os.Getenv("PATH")),
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}

	c = exec.Command("git", "config", "--global", "user.email", githubEmail)
	c.Env = []string{
		fmt.Sprintf("PATH=%s", os.Getenv("PATH")),
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

// isPROpen checks if a regen PR is already open.
func isPROpen(ctx context.Context, accessToken, username string) (bool, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	opt := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 50},
		State:       "open",
	}
	prs, _, err := githubClient.PullRequests.List(ctx, owner, repo, opt)
	if err != nil {
		return false, err
	}
	for _, pr := range prs {
		if !strings.Contains(pr.GetTitle(), "auto-regenerate") {
			continue
		}
		if pr.GetUser().GetLogin() != username {
			continue
		}
		return true, nil
	}
	return false, nil
}

// generate regenerates the whole project.
func generate(dir, remoteCacheDir string) error {
	oldCache := os.Getenv(remoteCacheEnv)
	os.Setenv(remoteCacheEnv, remoteCacheDir)
	defer func() {
		os.Setenv(remoteCacheEnv, oldCache)
	}()
	fp := filepath.Join(dir, "google-api-go-generator")
	cmd := exec.Command("make", "all")
	cmd.Dir = fp
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// hasChanges reports if any files have been updated.
func hasChanges(dir string) (bool, error) {
	c := exec.Command("git", "status", "--short")
	c.Dir = dir
	b, err := c.Output()
	return len(b) > 0, err
}

// makePR commits local changes and makes a regen PR.
func makePR(ctx context.Context, accessToken, dir string, dryRun bool) error {
	if dryRun {
		log.Println("dry run, not making PR")
		return nil
	}

	log.Println("creating commit and pushing")
	c := exec.Command("/bin/bash", "-c", `
	set -ex
	
	git config credential.helper store
	
	git branch -D $BRANCH_NAME || true
	git push -d origin $BRANCH_NAME || true
	
	git add -A
	git checkout -b $BRANCH_NAME
	git commit -m "$COMMIT_TITLE"
	git push origin $BRANCH_NAME
	`)
	c.Env = []string{
		fmt.Sprintf("COMMIT_TITLE=%s", commitTitle),
		fmt.Sprintf("BRANCH_NAME=%s", branchName),
		fmt.Sprintf("PATH=%s", os.Getenv("PATH")),
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}
	c.Dir = dir
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	if err := c.Run(); err != nil {
		return err
	}

	log.Println("creating pull request")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	head := owner + ":" + branchName
	base := "main"
	t := commitTitle
	_, _, err := githubClient.PullRequests.Create(ctx, owner, repo, &github.NewPullRequest{
		Title: &t,
		Head:  &head,
		Base:  &base,
	})
	if err != nil {
		return err
	}
	return nil
}

// cloneDiscoArtifactManager returns the directory of discovery documents found
// in the cloned repo.
func cloneDiscoArtifactManager() (string, error) {
	tmpDir, err := os.MkdirTemp("", "discogen")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:      "https://github.com/googleapis/discovery-artifact-manager.git",
		Progress: os.Stdout,
		Depth:    1,
		Tags:     git.NoTags,
	}); err != nil {
		return "", err
	}
	return filepath.Join(tmpDir, "discoveries"), nil
}
