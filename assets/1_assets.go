package assets

import (
	"embed"
)

// Embed assets folder inside the executable
//go:embed *
var FS embed.FS
