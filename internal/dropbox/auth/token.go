package auth

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"golang.org/x/oauth2"
)

func Token() {

}

func WriteToken(token *oauth2.Token) {

	jsonString, _ := json.Marshal(token)
	err := os.WriteFile("token.json", jsonString, os.ModePerm)
	if err != nil {
		log.Fatal("Error saving authentication information.")
	}
}

func ReadToken() (*oauth2.Token, error) {
	dat, err := os.ReadFile("token.json")
	if err != nil {
		return nil, errors.New("no token found")
	}
	if dat != nil {
		var token oauth2.Token
		json.Unmarshal(dat, &token)
		return &token, nil
	}
	return nil, errors.New("no data")
}
