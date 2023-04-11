package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
)

func TestListRepos(t *testing.T) {

	appID := int64(12345)
	installationID := int64(67890)
	privateKey := os.Getenv("PRIVATE_KEY")

	// Create a test server to mock the GitHub API responses
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name":"test-repo"}]`))
	}))
	defer testServer.Close()

	// Use the test server URL for the GitHub API endpoint
	client := github.NewClient(&http.Client{})
	client.BaseURL = testServer.URL + "/"

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.New(http.DefaultTransport, appID, installationID, []byte(privateKey))
	if err != nil {
		t.Fatal(err)
	}
	client.Transport = itr

	repos, _, err := client.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the list of repositories contains the expected repository name
	assert.Equal(t, "test-repo", *repos[0].Name)
}
