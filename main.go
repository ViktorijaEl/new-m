package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func main() {

	appID, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	if err != nil {
		log.Fatal("failed to parse APP_ID: ", err)
	}

	installationID, err := strconv.ParseInt(os.Getenv("INSTALLATION_ID"), 10, 64)
	if err != nil {
		log.Fatal("failed to parse INSTALLATION_ID: ", err)
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("PRIVATE_KEY environment variable is not set")
	}

	fmt.Println("PRIVATE_KEY: ", privateKey)

	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.NewKeyFromFile(tr, appID, installationID, privateKey)

	if err != nil {
		log.Fatal("failed to create ghinstallation transport: ", err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	repos, _, err := client.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		log.Fatal("failed to list repos: ", err)
	}

	// Print the name of each repository
	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}


