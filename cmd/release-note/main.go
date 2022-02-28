package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/google/go-github/v42/github"
)

const (
	owner  = "micro"
	repo   = "services"
	branch = "master"
)

func main() {
	ctx := context.Background()
	// create token source
	// ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: secert})
	// // create oauth http client
	// tc := oauth2.NewClient(ctx, ts)
	// client := github.NewClient(tc)
	client := github.NewClient(nil)
	branch, _, err := client.Repositories.GetBranch(ctx, owner, repo, branch, false)
	if err != nil {
		log.Fatal(err)
	}

	commits := branch.GetCommit()

	// extracting html_url, sha and message
	htm_url := commits.GetHTMLURL()
	sha := commits.GetSHA()
	message := commits.GetCommit().GetMessage()

	payload := map[string]interface{}{
		"sha":     sha[:9],
		"url":     htm_url,
		"message": message,
	}

	release_note_temp := `[{{ .sha }}]({{ .url }}) {{ .message }}`

	// create a template for release note
	t, err := template.New("release_note").Parse(release_note_temp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "faild to parse %v - err: %v\n", release_note_temp, err)
	}

	var tb bytes.Buffer
	err = t.Execute(&tb, payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "faild to apply parsed template %s to payload %v - err: %v\n", release_note_temp, payload, err)
	}

	fmt.Fprint(os.Stdout, tb.String())
}
