package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v29/github"
	"github.com/sethvargo/go-githubactions"
)

type PullRequestEvent struct {
	PullRequest `json:"pull_request"`
	Repository  `json:"repository"`
}

type PullRequest struct {
	Number int
}

type Repository struct {
	Name  string
	Owner `json:"owner"`
}

type Owner struct {
	Login string
}

func main() {

	githubToken := githubactions.GetInput("GITHUB_TOKEN")
	if githubToken == "" {
		githubactions.Fatalf("missing input 'githubToken'")
	}

	event, err := ioutil.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		log.Printf("Can't read events: " + err.Error())
	}

	var pr PullRequestEvent
	json.Unmarshal(event, &pr)

	log.Printf("Number: %d Repo Name: %s, Owner: %s", pr.PullRequest.Number, pr.Repository.Name, pr.Owner.Login)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	message := "Hello World"
	comment := &github.IssueComment{
		Body: &message,
	}
	_, _, err = client.Issues.CreateComment(ctx, pr.Owner.Login, pr.Repository.Name, pr.Number, comment)
	if err != nil {
		log.Printf(err.Error())
	}
}
