package main

import (
	"fmt"
	"log"

	audiofiles "github.com/cuthbeorht/media-library-manager/internal/audio_files"
	"github.com/cuthbeorht/media-library-manager/internal/config"
	database "github.com/cuthbeorht/media-library-manager/internal/database/sql"
	"github.com/cuthbeorht/media-library-manager/internal/dropbox/auth"
	"github.com/cuthbeorht/media-library-manager/internal/dropbox/files"
	walker "github.com/cuthbeorht/media-library-manager/pkg/library"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Initializing Media Library Manager")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not find environment variavles")
	}

	appConfig := config.NewApplicationConfig()

	dbConn := database.Connect(appConfig.DatabasePath)
	defer dbConn.Close()
	audiofiles.CreateTables(dbConn)

	token := auth.Auth(appConfig)
	dropboxConfig := config.DropboxConfig(token.AccessToken)
	filesClient := files.Client(dropboxConfig)

	walker.WalkMediaDir(dbConn, filesClient, "/New_Music/music")

}
