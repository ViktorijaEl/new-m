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
	"github.com/joho/godotenv"
)

func main() {

	// for local development, if statement checks if WSL_DISTRO_NAME=Ubuntu, then load env vars from .env file
	if os.Getenv("WSL_DISTRO_NAME") == "Ubuntu" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file")
			return
		}
	}

	appID, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	installationID, err := strconv.ParseInt(os.Getenv("INSTALLATION_ID"), 10, 64)
//     	privateKey := "./list-mg-repos.2023-04-05.private-key.pem"
// 	privateKey := "private.pem"
	privateKey := os.Getenv("PRIVATE_KEY")
	fmt.Println("PRIVATE_KEY: ", privateKey)
	fmt.Println("Private key contents:", privateKey)
// 	fmt.Println(privateKey)
	
	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
// 	itr, err := ghinstallation.NewKeyFromFile(tr, appID, installationID, privateKey)
	itr, err := ghinstallation.New(tr, appID, installationID, []byte(privateKey))

	if err != nil {
		log.Fatal(err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	repos, _, err := client.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	// fmt.Println(repos[0])

	// Print the name of each repository
	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}

