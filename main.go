package main

import (
	"log"
	"os"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	githubToken := githubactions.GetInput("GITHUB_TOKEN")
	if githubToken == "" {
		githubactions.Fatalf("missing input 'githubToken'")
	}

	event := os.Getenv("GITHUB_EVENT_PATH")

	log.Printf(event)

	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource(
	// 	&oauth2.Token{AccessToken: githubToken},
	// )
	// tc := oauth2.NewClient(ctx, ts)
	// client := github.NewClient(tc)

	// client.Issues.CreateComment(ctx, )

}
