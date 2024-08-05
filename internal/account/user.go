package user

import (
	"fmt"
	"log"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/users"
)

func CurrentUser(client users.Client) *users.FullAccount {
	basincInfo, err := client.GetCurrentAccount()
	if err != nil {
		fmt.Println("Could not get account id")
	}

	return basincInfo
}

func UserAccount(client users.Client, accountId string) *users.BasicAccount {
	accountArgs := users.GetAccountArg{AccountId: accountId}
	user, err := client.GetAccount(&accountArgs)
	if err != nil {
		fmt.Println("Could not get user: ", err)
		log.Fatal(err)
	}

	return user
}
