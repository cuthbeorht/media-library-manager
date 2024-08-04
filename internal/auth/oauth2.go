package auth

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
)

func Auth() *oauth2.Token {

	token, _ := ReadToken()
	if token != nil {
		return token
	}

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "uh8fq8z3ah7tlex",
		ClientSecret: "xj75sln1z8e23jn",
		Scopes:       []string{"files.metadata.read", "account_info.read", "sharing.read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.dropbox.com/oauth2/authorize",
			TokenURL: "https://api.dropboxapi.com/oauth2/token",
		},
	}

	verifier := oauth2.GenerateVerifier()

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	fmt.Print("\nEnter the confirmation code: ")
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}
	WriteToken(tok)

	return tok

}
