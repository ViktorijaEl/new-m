package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/google/go-github/v33/github"
)

func TestGetRepositories(t *testing.T) {
    // Create a test server with a mocked response
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, `[{"name": "repo1"}, {"name": "repo2"}]`)
    }))
    defer ts.Close()

    // Set up a GitHub client that uses the test server as its base URL
    client := github.NewClient(nil)
    baseURL := ts.URL + "/"
    u, err := url.Parse(baseURL)
    if err != nil {
        t.Fatalf("error parsing URL: %v", err)
    }
    client.BaseURL = u

    // Call the function being tested
    repos, err := GetRepositories(client)
    if err != nil {
        t.Fatalf("error getting repositories: %v", err)
    }

    // Verify the results
    if len(repos) != 2 {
        t.Fatalf("expected 2 repositories, got %d", len(repos))
    }
    if repos[0].Name != "repo1" {
        t.Errorf("expected first repository name to be repo1, got %s", repos[0].Name)
    }
    if repos[1].Name != "repo2" {
        t.Errorf("expected second repository name to be repo2, got %s", repos[1].Name)
    }
}
