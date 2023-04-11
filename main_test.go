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
	// Set up a test server to receive the HTTP requests from the client
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"name": "repo1"}, {"name": "repo2"}]`))
	}))
	defer testServer.Close()

	// Use the test server URL as the base URL for the client
	baseURL := testServer.URL + "/"

	// Set up a client with a transport that uses the test server URL
	tr := http.DefaultTransport
	itr, err := ghinstallation.New(tr, 1, 99, []byte(os.Getenv("PRIVATE_KEY")))
	if err != nil {
		t.Fatal(err)
	}
	client := github.NewClient(&http.Client{Transport: itr})
	client.BaseURL = baseURL

	// Call the ListRepos method on the client
	repos, _, err := client.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the returned repositories have the expected names
	if len(repos) != 2 {
		t.Fatalf("expected 2 repositories, got %d", len(repos))
	}
	if *repos[0].Name != "repo1" {
		t.Errorf("expected first repository name to be 'repo1', got '%s'", *repos[0].Name)
	}
	if *repos[1].Name != "repo2" {
		t.Errorf("expected second repository name to be 'repo2', got '%s'", *repos[1].Name)
	}
}
