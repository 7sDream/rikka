package fs

import (
	"flag"

	"github.com/7sDream/rikka/plugins"
)

// plugin type
type fsPlugin struct{}

var (
	l = plugins.SubLogger("[FS]")

	argFilesDir     = flag.String("dir", "files", "Where files will be save when use fs plugin.")
	argFsDebugSleep = flag.Int("fsDebugSleep", 0, "Debug: sleep some ms before copy file to fs, used to test javascript ajax")

	imageDir string

	// Plugin is the main plugin instance.
	Plugin = fsPlugin{}
)
