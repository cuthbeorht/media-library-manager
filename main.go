package main

import (
	"fmt"

	manager "github.com/cuthbeorht/media-library-manager/internal"
)

func main() {
	fmt.Println("Initializing Media Library Manager")

	manager.WalkMediaDir("/Volumes/DropboxData/Music")

}
