package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func TestGetRepositories(t *testing.T) {
	// Set up a test server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer testServer.Close()

	// Use the test server's URL in place of the GitHub API's URL
	baseURL := testServer.URL

	// Set up a new http client with the test server's URL and authentication
	tr := http.DefaultTransport
	itr, err := ghinstallation.New(tr, 1234, 5678, []byte("dummy key"))
	if err != nil {
		t.Errorf("Failed to create installation transport: %v", err)
	}
	client := github.NewClient(&http.Client{Transport: itr, BaseURL: baseURL})

	// Call the function to get the repositories
	repos, err := client.Repositories.List(nil)
	if err != nil {
		t.Errorf("Failed to get repositories: %v", err)
	}

	// Ensure that the returned repository count is 0
	if len(repos) != 0 {
		t.Errorf("Expected 0 repositories, but got %d", len(repos))
	}
}
