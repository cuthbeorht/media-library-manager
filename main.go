package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		mode     string //WEB or TTY
		rootPath string //Starting point from where we start parsing
	)

	flag.StringVar(&mode, "mode", "TTY", "Execution mode (TTY or WEB)")
	flag.StringVar(&rootPath, "root-path", ".", "Root path")
	flag.Parse()

	fmt.Printf("Starting application via TTY in %s mode with a path of %s", mode, rootPath)

	// cli.Cli()
}
