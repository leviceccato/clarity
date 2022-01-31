package main

import (
	"fmt"
	"os"
	"os/exec"
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
	err := exec.Command("go", "build").Run()
	if err != nil {
		panic(fmt.Sprintf("building macos executable: %s", err))
	}

	for _, folder := range []string{
		"Clarity.app",
		"Clarity.app/Contents",
		"Clarity.app/Contents/MacOS",
		"Clarity.app/Contents/Resources",
	} {
		err := os.Mkdir(folder, 0755)
		if err != nil {
			panic(fmt.Sprintf("creating %s folder: %s", folder, err))
		}
	}
	err = os.WriteFile("Clarity.app/Contents/Info.plist", []byte(infoTmpl), 0777)
	if err != nil {
		panic(fmt.Sprintf("creating Info.plist: %s", err))
	}
	bin, err := os.ReadFile("clarity")
	if err != nil {
		panic(fmt.Sprintf("copying clarity binary: %s", err))
	}
	err = os.WriteFile("Clarity.app/Contents/MacOS/Clarity", bin, 0777)
	if err != nil {
		panic(fmt.Sprintf("pasting clarity binary: %s", err))
	}

	// Run macOS specific iconutil command to generate icons
	err = exec.Command(
		"iconutil", "-c", "icns", "-o",
		"Clarity.app/Contents/Resources/icon.icns", "asset/icon.iconset",
	).Run()
	if err != nil {
		panic(fmt.Sprintf("generating icons with iconutil: %s", err))
	}
}
