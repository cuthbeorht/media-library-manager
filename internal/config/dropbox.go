package config

import "github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"

func DropboxConfig(accessToken string) dropbox.Config {
	return dropbox.Config{
		Token:    accessToken,
		LogLevel: dropbox.LogDebug,
	}
}
