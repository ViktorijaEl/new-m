package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
)

func TestGetRepos(t *testing.T) {
	// Create a test HTTP server and a client that will use it
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[{"name":"test-repo-1"},{"name":"test-repo-2"}]`))
	}))
	defer testServer.Close()

	client := github.NewClient(nil)
	// Override client's base URL with the test server's URL
	baseURL := testServer.URL + "/"
	client.BaseURL, _ = url.Parse(baseURL)

	// Call the function being tested
	repos, err := getRepos(client)
	if err != nil {
		t.Errorf("Error getting repositories: %v", err)
	}

	// Verify the result
	if len(repos) != 2 {
		t.Errorf("Expected 2 repositories, but got %d", len(repos))
	}
	if repos[0].Name != "test-repo-1" || repos[1].Name != "test-repo-2" {
		t.Errorf("Unexpected repository names: %v", repos)
	}
}
