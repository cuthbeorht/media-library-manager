package manager

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkMediaDir(root string) {
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		fmt.Println("Filename: ", info.Name())

		return nil
	})
}
