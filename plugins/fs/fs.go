package fs

import (
	"flag"

	"github.com/7sDream/rikka/plugins"
)

var l = plugins.SubLogger("[FS]")

var argFilesDir = flag.String("dir", "files", "Where files will be save when use fs plugin.")
var argFsDebugSleep = flag.Int("fsDebugSleep", 0, "Debug: sleep some ms before copy file to fs, used to test javascripta ajax")

var imageDir string

// plugin type
type fsPlugin struct{}

// FsPlugin is the main plugin instance.
var FsPlugin = fsPlugin{}
