package main

import (
	"fmt"

	"github.com/cuthbeorht/media-library-manager/internal/config"
	"github.com/cuthbeorht/media-library-manager/internal/dropbox/auth"
	"github.com/cuthbeorht/media-library-manager/internal/dropbox/files"
	"github.com/cuthbeorht/media-library-manager/internal/walker"
)

func main() {
	fmt.Println("Initializing Media Library Manager")

	appConfig := config.NewApplicationConfig()
	token := auth.Auth(appConfig)
	dropboxConfig := config.DropboxConfig(token.AccessToken)
	filesClient := files.Client(dropboxConfig)

	walker.WalkMediaDir(filesClient, "/New_Music/music/Mass Effect 2 OST")

}
