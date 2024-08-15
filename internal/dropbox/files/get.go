package files

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

func ListFolder(client files.Client) []string {
	var audioFiles []string

	args := files.NewListFolderArg("/Music")
	client.ListFolder(args)

	return audioFiles
}
