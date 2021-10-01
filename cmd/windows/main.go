package main

import (
	"fmt"
	"image"
	"os"

	"github.com/leviceccato/clarity/util"

	"github.com/tc-hib/winres"
)

// This package uses tc-hib/winres to create a resource set that gets embedded
// into the final executable for Windows. This allows embedding the application
// icon. This file should not be included in the build and instead should be
// run with `go run ./windows`.
func main() {
	rs := winres.ResourceSet{}

	// Add icons
	icon32, err := util.LoadIcon("icon.iconset/icon_32x32.png")
	if err != nil {
		fmt.Printf("loading icon_32: %s", err)
		return
	}
	icon16, err := util.LoadIcon("icon.iconset/icon_16x16.png")
	if err != nil {
		fmt.Printf("loading icon_16: %s", err)
		return
	}
	icon, err := winres.NewIconFromImages([]image.Image{icon32, icon16})
	if err != nil {
		fmt.Printf("creating icon from files: %s", err)
		return
	}
	rs.SetIcon(winres.Name("APPICON"), icon)

	// Output syso file for inclusion in executable
	out, err := os.Create("rsrc_windows_amd64.syso")
	if err != nil {
		fmt.Printf("creating syso file: %s", err)
		return
	}
	defer out.Close()
	err = rs.WriteObject(out, winres.ArchAMD64)
	if err != nil {
		fmt.Printf("writing resource to syso file: %s", err)
		return
	}
}
