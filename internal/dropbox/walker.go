package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/users"
)

func WalkMediaDir(token string) {
	config := dropbox.Config{
		Token:    token,
		LogLevel: dropbox.LogDebug,
	}

	dbx := users.New(config)
	dbx.GetAccount(&users.GetAccountArg{AccountId: ""})
}
