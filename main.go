package main

import (
	"flag"
	"fmt"

	"github.com/cuthbeorht/media-library-manager/cmd/cli"
)

func main() {
	var (
		mode     string //WEB or TTY
		sources  string //Sources for media
		rootPath string //Starting point from where we start parsing
	)

	flag.StringVar(&mode, "mode", "TTY", "Execution mode (TTY or WEB)")
	flag.StringVar(&rootPath, "root-path", ".", "Root path")
	flag.StringVar(&sources, "sources", "DROPBOX", "Souces of media files. Comma delimited value.")
	flag.Parse()

	fmt.Printf("Starting application via TTY in %s mode with a path of %s", mode, rootPath)

	if mode == "TTY" && sources == "DROPBOX" {
		cli.Cli()
	}
}
