package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func TestMain(t *testing.T) {
	// Set up test environment variables
	os.Setenv("APP_ID", "123")
	os.Setenv("INSTALLATION_ID", "456")
	os.Setenv("PRIVATE_KEY", "test-private-key")

	// Create test client
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	testServer := httptest.NewServer(handler)

	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(testServer.URL)
		},
	}
	defer testServer.Close()

	// Create test context and transport
	ctx := context.Background()
	itr, err := ghinstallation.New(tr, 123, 456, []byte("test-private-key"))
	if err != nil {
		t.Fatalf("Error creating installation transport: %v", err)
	}
	client := github.NewClient(&http.Client{Transport: itr})

	// Create test repositories
	repos := []*github.Repository{
		{Name: github.String("test-repo-1")},
		{Name: github.String("test-repo-2")},
		{Name: github.String("test-repo-3")},
	}

	// Set up test response
	respBody := bytes.NewBuffer(nil)
	for _, repo := range repos {
		respBody.WriteString(*repo.Name + "\n")
	}
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(respBody),
	}

	// Test main function
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)

	if string(out) != "test-repo-1\ntest-repo-2\ntest-repo-3\n" {
		t.Errorf("Unexpected output: %v", string(out))
	}

	// Test client method call
	client.Transport = tr
	client.BaseURL = testServer.URL
	gotRepos, _, err := client.Apps.ListRepos(ctx, nil)
	if err != nil {
		log.Fatalf("Error listing repositories: %v", err)
	}

	if len(gotRepos) != len(repos) {
		t.Errorf("Expected %d repositories, but got %d", len(repos), len(gotRepos))
	}
	for i := range repos {
		if *gotRepos[i].Name != *repos[i].Name {
			t.Errorf("Expected repository name %q, but got %q", *repos[i].Name, *gotRepos[i].Name)
		}
	}
}
