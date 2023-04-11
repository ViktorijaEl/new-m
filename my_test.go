package main

import (
	"context"
	"os"
	"strconv"
	"testing"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func TestAuthorization(t *testing.T) {
	appID, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	if err != nil {
		t.Fatalf("Failed to parse APP_ID: %v", err)
	}
	installationID, err := strconv.ParseInt(os.Getenv("INSTALLATION_ID"), 10, 64)
	if err != nil {
		t.Fatalf("Failed to parse INSTALLATION_ID: %v", err)
	}
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		t.Fatal("PRIVATE_KEY is not set")
	}

	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.New(tr, appID, installationID, []byte(privateKey))
	if err != nil {
		t.Fatalf("Failed to create installation transport: %v", err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	_, _, err = client.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		t.Fatalf("Authorization failed: %v", err)
	}
}
