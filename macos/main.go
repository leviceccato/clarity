package main

import (
	"fmt"
	"os"
)

var folders = []string{
	"clarity.app",
	"clarity.app/Contents",
	"clarity.app/Contents/MacOS",
	"clarity.app/Contents/Resources",
}

// This package uses create the folder structure and config for a macOS .app
// folder. This will allow you to have an embedded icon.
func main() {
	for _, folder := range folders {
		if err := os.Mkdir(folder, 0755); err != nil {
			fmt.Printf("creating %s folder: %s", folder, err)
			return
		}
	}
}
