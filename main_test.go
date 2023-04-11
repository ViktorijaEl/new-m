package main

import (
	"context"
	"net/http"
    "net/url"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
)

func TestListRepos(t *testing.T) {
	// Set up a test server with a sample response
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assert that the correct URL is called
		assert.Equal(t, "/app/installations/1/repositories", r.URL.String())

		// Return a sample response
		w.Write([]byte(`[{"name":"repo1"},{"name":"repo2"}]`))
	}))
	defer testServer.Close()

	// Use the test server URL as the API base URL
	client := github.NewClient(nil)
	client.BaseURL = &testServer.URL

	// Make the API request and assert that it returns the expected response
	repos, _, err := client.Apps.ListRepos(context.Background(), nil)
	assert.NoError(t, err)
	assert.Len(t, repos, 2)
	assert.Equal(t, "repo1", *repos[0].Name)
	assert.Equal(t, "repo2", *repos[1].Name)
}
