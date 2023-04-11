package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func TestMain(m *testing.M) {
	// Set up a test server to respond to API requests
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[{"name": "repo1"}, {"name": "repo2"}]`))
	}))
	defer testServer.Close()

	// Override the GitHub API endpoint to use the test server
	os.Setenv("GITHUB_API_URL", testServer.URL)

	// Run the tests
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestListRepos(t *testing.T) {
	appID, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	installationID, err := strconv.ParseInt(os.Getenv("INSTALLATION_ID"), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	privateKey := os.Getenv("PRIVATE_KEY")

	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.New(tr, appID, installationID, []byte(privateKey))
	if err != nil {
		t.Fatal(err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	repos, _, err := client.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(repos) != 2 {
		t.Fatalf("Expected 2 repos, but got %d", len(repos))
	}

	if *repos[0].Name != "repo1" {
		t.Errorf("Expected first repo name to be repo1, but got %s", *repos[0].Name)
	}

	if *repos[1].Name != "repo2" {
		t.Errorf("Expected second repo name to be repo2, but got %s", *repos[1].Name)
	}
}
