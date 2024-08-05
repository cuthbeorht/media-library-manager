package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/users"
)

func Client(accessToken string) users.Client {
	config := dropbox.Config{
		Token:    accessToken,
		LogLevel: dropbox.LogDebug,
	}

	return users.New(config)
}
