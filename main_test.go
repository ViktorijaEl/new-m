package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func main() {
	appID, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	if err != nil {
		log.Fatalf("error parsing APP_ID: %s", err)
	}

	installationID, err := strconv.ParseInt(os.Getenv("INSTALLATION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("error parsing INSTALLATION_ID: %s", err)
	}

	privateKey := os.Getenv("PRIVATE_KEY")

	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.New(tr, appID, installationID, []byte(privateKey))
	if err != nil {
		log.Fatalf("error creating ghinstallation transport: %s", err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	// List the repositories for the authenticated installation
	ctx := context.Background()
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatalf("error listing repositories: %s", err)
	}

	// Print the name of each repository
	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}
