//go:build !embed

package gfx

import (
	"io/fs"
	"os"
)

var assetsFS fs.FS

func init() {
	if _, err := os.Stat("../assets"); err == nil {
		assetsFS = os.DirFS("../assets/PNG/Dark")
	} else {
		assetsFS = os.DirFS("assets/PNG/Dark")
	}
}
