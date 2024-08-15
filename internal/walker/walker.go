package walker

import (
	"database/sql"
	"fmt"
	"log"

	dropbox "github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

func WalkMediaDir(dbConn *sql.DB, client dropbox.Client, path string) {

	var getFiles func(client dropbox.Client, path string) []string

	getFiles = func(client dropbox.Client, path string) []string {
		content := getPathContent(client, path)
		var myFiles []string

		for _, entry := range content.Entries {
			switch f := entry.(type) {
			case *dropbox.FolderMetadata:
				// fmt.Println("Folder entry: ", f.Name, f.PathDisplay)
				newFiles := getFiles(client, f.PathDisplay)
				myFiles = append(myFiles, newFiles...)

			case *dropbox.FileMetadata:
				// fmt.Println("File content: ", f.Name)
				myFiles = append(myFiles, f.Name)
			}

		}

		return myFiles
	}

	myFiles := getFiles(client, path)

	for _, x := range myFiles {

		if fileName == "" {

			fmt.Println("File: ", x, "\nInsert result: ", res)
		}

	}

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
