// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2024-11-23 20:25:15

// === Go Playground ===
// Execute the snippet with:                 Ctl-Return
// Provide custom arguments to compile with: Alt-Return
// Other useful commands:
// - remove the snippet completely with its dir and all files: (go-playground-rm)
// - upload the current buffer to playground.golang.org:       (go-playground-upload)

package main

import (
	"log"
	"os"
)

func main() {
	app := NewMainApp()
	err := RunMainApp(app, os.Args...)
	if err != nil {
		log.Fatal(err)
	}
}
