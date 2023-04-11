package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
)

func TestListRepos(t *testing.T) {
	// Create a test server with a mocked response
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the correct API endpoint is being called
		if r.URL.Path != "/app/installations/123/repositories" {
			t.Errorf("unexpected request path: %s", r.URL.Path)
		}
		// Return a JSON response with one repository
		w.Write([]byte(`[{"id": 1, "name": "repo1"}]`))
	}))
	defer testServer.Close()

	// Create a new GitHub client that uses the test server as the HTTP transport
	client := github.NewClient(nil)
	client.BaseURL = testServer.URL + "/"

	// Set the client to use the test server's transport
	client.SetDo(func(req *http.Request) (*http.Response, error) {
		req.URL.Scheme = testServer.URL
		req.URL.Host = ""
		return http.DefaultTransport.RoundTrip(req)
	})

	// Call the ListRepos function and check the result
	repos, _, err := client.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(repos) != 1 || *repos[0].Name != "repo1" {
		t.Errorf("unexpected repos: %v", repos)
	}
}
