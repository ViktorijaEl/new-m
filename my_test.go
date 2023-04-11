package main

import (
	"context"
	"strconv"
	"testing"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func TestRepositoriesList(t *testing.T) {
	appID, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	if err != nil {
		t.Fatalf("error parsing APP_ID: %s", err)
	}

	installationID, err := strconv.ParseInt(os.Getenv("INSTALLATION_ID"), 10, 64)
	if err != nil {
		t.Fatalf("error parsing INSTALLATION_ID: %s", err)
	}

	privateKey := os.Getenv("PRIVATE_KEY")

	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.New(tr, appID, installationID, []byte(privateKey))
	if err != nil {
		t.Fatalf("error creating ghinstallation transport: %s", err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	// List the repositories for the authenticated installation
	ctx := context.Background()
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		t.Fatalf("error listing repositories: %s", err)
	}

	// Check that the repositories list is not empty
	if len(repos) == 0 {
		t.Error("expected at least one repository, but got none")
	}

	// Add more test cases here as needed
}
