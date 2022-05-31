package main

import (
	"os"
	"os/exec"

	"github.com/leviceccato/clarity/util"
)

const infoTmpl = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>Clarity</string>
	<key>CFBundleIconFile</key>
	<string>icon.icns</string>
	<key>CFBundleIdentifier</key>
	<string>com.clarity.game</string>
	<key>NSHighResolutionCapable</key>
	<true/>
</dict>
</plist>
`

// This package creates a folder structure and config for a macOS .app.
// This will allow you to have an embedded icon.
func main() {

	// Build executable
	util.Must(exec.Command("go", "build").Run())

	// Create .app folder and subfolders
	for _, folder := range []string{
		"Clarity.app",
		"Clarity.app/Contents",
		"Clarity.app/Contents/MacOS",
		"Clarity.app/Contents/Resources",
	} {
		util.Must(os.Mkdir(folder, 0755))
	}

	// Create .plist from template
	util.Must(os.WriteFile("Clarity.app/Contents/Info.plist", []byte(infoTmpl), 0777))

	// Run macOS specific iconutil command to generate icons
	util.Must(exec.Command(
		"iconutil", "-c", "icns", "-o",
		"Clarity.app/Contents/Resources/icon.icns", "asset/icon.iconset",
	).Run())

	// Copy binary into .app
	bin := util.MustGet(os.ReadFile("clarity"))
	util.Must(os.WriteFile("Clarity.app/Contents/MacOS/Clarity", bin, 0777))
}
