package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
)

func TestGetRepoName(t *testing.T) {
	// Create a test HTTP server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name": "repo1"}, {"name": "repo2"}]`))
	}))

	// Make sure the test server is closed when the test finishes
	defer testServer.Close()

	// Parse the test server URL
	baseURL, err := url.Parse(testServer.URL)
	if err != nil {
		t.Fatalf("Failed to parse test server URL: %v", err)
	}

	// Create a new GitHub client using the test server URL
	client := github.NewClient(nil)
	client.BaseURL = baseURL

	// Call the function being tested
	repoNames, err := GetRepoNames(client)
	if err != nil {
		t.Fatalf("Failed to get repo names: %v", err)
	}

	// Check the results
	expected := []string{"repo1", "repo2"}
	if !reflect.DeepEqual(repoNames, expected) {
		t.Errorf("Expected repo names %v, but got %v", expected, repoNames)
	}
}
