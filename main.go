package main

import (
	"fmt"

	"github.com/cuthbeorht/media-library-manager/internal/auth"
	"github.com/cuthbeorht/media-library-manager/internal/dropbox"
)

func main() {
	fmt.Println("Initializing Media Library Manager")

	token := auth.Auth()

	dropbox.WalkMediaDir(token.AccessToken)

}
