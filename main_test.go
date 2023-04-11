package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func TestListRepos(t *testing.T) {
	// Setup test server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer testServer.Close()

	// Use test server URL as base URL for client
	baseURL := testServer.URL + "/"

	// Create client
	tr := http.DefaultTransport
	itr, err := ghinstallation.New(tr, 1, 99, []byte(os.Getenv("PRIVATE_KEY")))
	if err != nil {
		t.Fatalf("Failed to create installation transport: %v", err)
	}
	client := github.NewClient(&http.Client{Transport: itr})
	client.BaseURL = baseURL

	// Make request
	repos, _, err := client.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		t.Fatalf("Failed to list repos: %v", err)
	}

	// Check response
	if len(repos) != 0 {
		t.Fatalf("Expected 0 repos, but got %d", len(repos))
	}
}
