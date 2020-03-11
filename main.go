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

// PullRequestEvent is the root of the event payload related to a pull request
// More details: https://developer.github.com/v3/activity/events/types/#pullrequestevent
type PullRequestEvent struct {
	PullRequest `json:"pull_request"`
	Repository  `json:"repository"`
}

// PullRequest contains the information describing the pr
type PullRequest struct {
	Number int
}

// Repository contains the information describing the repo the pr is beind made to
type Repository struct {
	Name  string
	Owner `json:"owner"`
}

// Owner contains the information about the org or individual who created the repo
type Owner struct {
	Login string
}

func main() {

	githubToken := githubactions.GetInput("GITHUB_TOKEN")
	if githubToken == "" {
		githubactions.Fatalf("missing input 'githubToken'")
		return
	}

	event, err := ioutil.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		log.Printf("Can't read events: " + err.Error())
		return
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

	// TODO make and add more Gopher gifs!
	url := "https://media.giphy.com/media/S6q921h64g8Qi305AJ/giphy.gif"

	message := "![GopherGif](" + url + ")"
	comment := &github.IssueComment{
		Body: &message,
	}
	_, _, err = client.Issues.CreateComment(ctx, pr.Owner.Login, pr.Repository.Name, pr.Number, comment)
	if err != nil {
		log.Printf(err.Error())
		return
	}
}
