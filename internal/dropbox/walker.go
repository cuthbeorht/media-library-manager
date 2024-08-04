package dropbox

import (
	"fmt"
	"log"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/users"
)

func WalkMediaDir(token string) {
	config := dropbox.Config{
		Token:    token,
		LogLevel: dropbox.LogDebug,
	}

	dbx := users.New(config)

	basincInfo, err := dbx.GetCurrentAccount()
	if err != nil {
		fmt.Println("Could not get account id")
	}
	fmt.Println("Account ID: ", basincInfo.AccountId)

	accountArgs := users.GetAccountArg{AccountId: basincInfo.AccountId}
	user, err := dbx.GetAccount(&accountArgs)
	if err != nil {
		fmt.Println("Could not get user: ", err)
		log.Fatal(err)
	}

	fmt.Println("User: ", user.Email)
}
