package main

import (
	"fmt"
	"image"
	"os"
	"os/exec"

	"github.com/leviceccato/clarity/asset"
	"github.com/tc-hib/winres"
)

// This package uses tc-hib/winres to create a resource set that gets embedded
// into the final executable for Windows. This allows embedding the application
// icon. This file should not be included in the build and instead should be
// run with `go run ./windows`.
func main() {
	rs := winres.ResourceSet{}

	// Add icons
	icon32, err := asset.LoadIcon("icon.iconset/icon_32x32.png")
	if err != nil {
		panic(fmt.Sprintf("loading icon_32: %s", err))
	}
	icon16, err := asset.LoadIcon("icon.iconset/icon_16x16.png")
	if err != nil {
		panic(fmt.Sprintf("loading icon_16: %s", err))
	}
	icon, err := winres.NewIconFromImages([]image.Image{icon32, icon16})
	if err != nil {
		panic(fmt.Sprintf("creating icon from files: %s", err))
	}
	rs.SetIcon(winres.Name("APPICON"), icon)

	// Output syso file for inclusion in executable
	out, err := os.Create("rsrc_windows_amd64.syso")
	if err != nil {
		panic(fmt.Sprintf("creating syso file: %s", err))
	}
	defer out.Close()
	err = rs.WriteObject(out, winres.ArchAMD64)
	if err != nil {
		panic(fmt.Sprintf("writing resource to syso file: %s", err))
	}

	// Build executable
	err = exec.Command("go", "build", "-o", "Clarity.exe").Run()
	if err != nil {
		panic(fmt.Sprintf("generating windows executable: %s", err))
	}
}
