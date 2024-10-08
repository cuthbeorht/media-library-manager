package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/cuthbeorht/media-library-manager/internal/config"
	"golang.org/x/oauth2"
)

func Auth(appConfig config.ApplicationConfig) *oauth2.Token {

	token, _ := ReadToken()
	if token != nil && token.Valid() {
		return token
	}

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     appConfig.ClientId,
		ClientSecret: appConfig.ClientSecret,
		Scopes:       appConfig.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  appConfig.AuthUrl,
			TokenURL: appConfig.TokenUrl,
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
