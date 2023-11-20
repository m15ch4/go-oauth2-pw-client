package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
)

func main() {
	// Configuring the OAuth2 client
	config := &oauth2.Config{
		ClientID:     "go-clien",
		ClientSecret: "blhZ1ueOSecw8batVFqHQEo2zts6WwdopXdYTfko62wX19Tu",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://xreg-vidm01.vmwpslab.pl/SAAS/auth/oauthtoken", // The token endpoint of the OAuth2 provider
		},
		Scopes: []string{"email", "profile", "user", "openid"}, // Specify the required scopes
	}

	// User credentials
	username := "configadmin"
	password := "VMw@re1!"

	// Performing the password credentials flow
	token, err := config.PasswordCredentialsToken(context.Background(), username, password)
	if err != nil {
		log.Fatalf("Error retrieving token: %v", err)
	}

	// Printing the obtained token
	fmt.Printf("Access Token: %s\n", token.AccessToken)
}
