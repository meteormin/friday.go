package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist
var assets embed.FS

var FS, _ = fs.Sub(assets, "dist")
