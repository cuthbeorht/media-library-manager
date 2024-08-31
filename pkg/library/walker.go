package walker

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/cuthbeorht/go-media-analyzer/src/media"

	audiofiles "github.com/cuthbeorht/media-library-manager/internal/audio_files"
	dropbox "github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

var validFileTypes = map[string]bool{
	"MP3":  true,
	"FLAC": true,
}

func WalkMediaDir(dbConn *sql.DB, client dropbox.Client, path string) {

	var getFiles func(client dropbox.Client, path string) []string

	getFiles = func(client dropbox.Client, path string) []string {
		content := getPathContent(client, path)
		var myFiles []string

		for _, entry := range content.Entries {
			switch f := entry.(type) {
			case *dropbox.FolderMetadata:

				getFiles(client, f.PathDisplay)

			case *dropbox.FileMetadata:
				newMedia, _ := media.OpenMedia(strings.Join([]string{"/Volumes/DropboxData", f.PathDisplay}, "/"))
				if validFileTypes[strings.ToUpper(newMedia.FileType)] {
					newFile := audiofiles.AudioFile{
						Name: newMedia.Name,
						Size: f.Size,
						Path: f.Metadata.PathDisplay,
					}
					audiofiles.PersistAudioFile(dbConn, newFile)
				}
			}

		}

		return myFiles
	}

	getFiles(client, path)

}

func getPathContent(client dropbox.Client, path string) *dropbox.ListFolderResult {
	fmt.Println("Fetching new batch from: ", path)
	myFiles, err := client.ListFolder(&dropbox.ListFolderArg{
		Path: path,
	})
	if err != nil {
		log.Fatal("unable to get directory: ", err)
	}
	return myFiles
}
