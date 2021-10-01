package main

import (
	"fmt"
	"os"
)

const infoTemplate = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>clarity</string>
	<key>CFBundleIconFile</key>
	<string>icon.icns</string>
	<key>CFBundleIdentifier</key>
	<string>com.example.yours</string>
	<key>NSHighResolutionCapable</key>
	<true/>
	<key>LSUIElement</key>
	<true/>
</dict>
</plist>
`

// This package creates a folder structure and config for a macOS .app.
// This will allow you to have an embedded icon.
func main() {
	for _, folder := range []string{
		"clarity.app",
		"clarity.app/Contents",
		"clarity.app/Contents/MacOS",
		"clarity.app/Contents/Resources",
	} {
		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Printf("creating %s folder: %s", folder, err)
			return
		}
	}
	err := os.WriteFile("clarity.app/Contents/Info.plist", []byte(infoTemplate), 0777)
	if err != nil {
		fmt.Printf("creating Info.plist: %s", err)
		return
	}
	binary, err := os.ReadFile("clarity")
	if err != nil {
		fmt.Printf("copying clarity binary: %s", err)
		return
	}
	err = os.WriteFile("clarity.app/Contents/MacOS/clarity", binary, 0777)
	if err != nil {
		fmt.Printf("pasting clarity binary: %s", err)
		return
	}
}
