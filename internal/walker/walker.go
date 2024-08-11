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
				getFiles(client, f.PathDisplay)

			case *dropbox.FileMetadata:
				// fmt.Println("File content: ", f.Name)
				myFiles = append(myFiles, f.Name)
			}

		}
		return myFiles
	}

	myFiles := getFiles(client, path)
	preparedInsert, _ := dbConn.Prepare("insert into audio_files (name) values (?)")
	for _, x := range myFiles {
		res, err := preparedInsert.Exec(x)
		if err != nil {
			fmt.Println("could not insert: ", err)
		}
		fmt.Println("File: ", x, "\nInsert result: ", res)
	}

}

func getPathContent(client dropbox.Client, path string) *dropbox.ListFolderResult {
	myFiles, err := client.ListFolder(&dropbox.ListFolderArg{
		Path: path,
	})
	if err != nil {
		log.Fatal("unable to get directory: ", err)
	}
	return myFiles
}
