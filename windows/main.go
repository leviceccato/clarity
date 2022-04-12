package main

import (
	"image"
	"os"
	"os/exec"

	"github.com/leviceccato/clarity/asset"
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
	icon32 := util.MustGet(asset.LoadIcon("icon.iconset/icon_32x32.png"))
	icon16 := util.MustGet(asset.LoadIcon("icon.iconset/icon_16x16.png"))
	icon := util.MustGet(winres.NewIconFromImages([]image.Image{icon32, icon16}))
	rs.SetIcon(winres.Name("APPICON"), icon)

	// Output syso file for inclusion in executable
	out := util.MustGet(os.Create("rsrc_windows_amd64.syso"))
	defer out.Close()
	util.Must(rs.WriteObject(out, winres.ArchAMD64))

	// Build executable
	util.Must(exec.Command("go", "build", "-o", "Clarity.exe").Run())
}
