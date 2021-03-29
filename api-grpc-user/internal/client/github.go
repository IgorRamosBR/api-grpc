package client

import (
	"github.com/google/go-github/v33/github"
)

func CreateGithubClient() *github.Client {
	return github.NewClient(nil)
}
