package main

import (
	"fmt"

	"github.com/cuthbeorht/media-library-manager/internal/auth"
	"github.com/cuthbeorht/media-library-manager/internal/config"
	"github.com/cuthbeorht/media-library-manager/internal/files"
	"github.com/cuthbeorht/media-library-manager/internal/walker"
)

func main() {
	fmt.Println("Initializing Media Library Manager")

	token := auth.Auth()
	appConfig := config.Config(token.AccessToken)
	filesClient := files.Client(appConfig)

	walker.WalkMediaDir(filesClient, "/New_Music/music/Mass Effect 2 OST")

}
