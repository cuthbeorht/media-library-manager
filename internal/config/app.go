package config

import (
	"os"
	"strings"
)

type ApplicationConfig struct {
	ClientId     string
	ClientSecret string
	Scopes       []string
	AuthUrl      string
	TokenUrl     string
}

func NewApplicationConfig() ApplicationConfig {
	return ApplicationConfig{
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       strings.Split(os.Getenv("SCOPES"), ","),
		AuthUrl:      os.Getenv("AUTH_URL"),
		TokenUrl:     os.Getenv("TOKEN_URL"),
	}
}
